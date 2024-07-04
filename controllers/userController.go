package controllers

import (
	"net/http"
	"strconv"

	"final-project-golang-individu/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetProfile godoc
// @Summary Get the profile of the logged-in user
// @Description Get the profile of the logged-in user
// @Tags user
// @Success 200 {object} models.Profile
// @Router /profile [get]
func (ctrl *UserController) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	user, err := ctrl.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, user.Profile)
}

// UpdateProfile godoc
// @Summary Update the profile of the logged-in user
// @Description Update the profile of the logged-in user
// @Tags user
// @Accept  json
// @Produce  json
// @Param profile body models.Profile true "Profile"
// @Success 200 {object} models.Profile
// @Router /profile [put]
func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	user, err := ctrl.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	if err := c.ShouldBindJSON(&user.Profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
// @Param password body string true "New Password"
// @Success 200 {string} string "Password updated successfully"
// @Router /change-password [put]
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		Password string `json:"password" binding:"required"`
	}
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
