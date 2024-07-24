package controllers

import (
	"final-project-golang-individu/repositories"
	"net/http"
	"strconv"

	"final-project-golang-individu/models"
	"final-project-golang-individu/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService    services.UserService
	roleRepository repositories.RoleRepository // Tambahkan RoleRepository
}

func NewUserController(userService services.UserService, roleRepository repositories.RoleRepository) *UserController {
	return &UserController{
		userService:    userService,
		roleRepository: roleRepository, // Inisialisasi RoleRepository
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.RegisterInput true "User Input"
// @Success 201 {object} models.User
// @Security BearerAuth
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get role by name
	role, err := ctrl.roleRepository.GetRoleByName(input.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Email:    input.Email,
	}

	// Create user
	if err := ctrl.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Assign role to user
	userRole := models.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}

	if err := ctrl.userService.AssignRoleToUser(&userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch the user with roles
	userWithRoles, err := ctrl.userService.GetUserWithRoles(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userWithRoles)
}

// UpdateUser godoc
// @Summary Update user details
// @Description Update the username, email, and role of a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.EditUserInput true "Edit User Input"
// @Success 200 {object} models.User
// @Security BearerAuth
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var input models.EditUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get existing user
	user, err := ctrl.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update fields
	user.Username = input.Username
	user.Email = input.Email

	// Get role by name
	role, err := ctrl.roleRepository.GetRoleByName(input.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})
		return
	}

	// Update user
	if err := ctrl.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove existing roles
	if err := ctrl.userService.RemoveRolesFromUser(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Assign new role
	userRole := models.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}

	if err := ctrl.userService.AssignRoleToUser(&userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch user with updated roles
	updatedUser, err := ctrl.userService.GetUserWithRoles(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user with roles"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// userController.go
// GetProfile godoc
// @Summary Get the profile of the logged-in user
// @Description Get the profile of the logged-in user
// @Tags user
// @Success 200 {object} models.User
// @Security BearerAuth
// @Router /profile [get]
func (ctrl *UserController) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	user, err := ctrl.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile godoc
// @Summary Update the profile of the logged-in user
// @Description Update the profile of the logged-in user
// @Tags user
// @Accept  json
// @Produce  json
// @Param profile body models.ProfileInput true "Profile Input"
// @Success 200 {object} models.Profile
// @Security BearerAuth
// @Router /profile [put]
func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var profileInput models.ProfileInput
	if err := c.ShouldBindJSON(&profileInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	user.Profile.Bio = profileInput.Bio
	user.Profile.Picture = profileInput.Picture

	if err := ctrl.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user.Profile)
}

// ChangePassword godoc
// @Summary Change the password of the logged-in user
// @Description Change the password of the logged-in user
// @Tags user
// @Accept  json
// @Produce  json
// @Param password body models.ChangePasswordInput true "Change Password Input"
// @Success 200 {string} string "Password updated successfully"
// @Security BearerAuth
// @Router /change-password [put]
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input models.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Password = string(hashedPassword)
	if err := ctrl.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Password updated successfully")
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Success 200 {array} models.User
// @Security BearerAuth
// @Router /users [get]
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags user
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Security BearerAuth
// @Router /users/{id} [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := ctrl.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags user
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Security BearerAuth
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = ctrl.userService.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
