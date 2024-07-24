package repositories

import (
	"final-project-golang-individu/models"
	"github.com/jinzhu/gorm"
)

// UserRepository interface
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	AssignRoleToUser(userRole *models.UserRole) error
	GetUserWithRoles(userID uint, user *models.User) error
	RemoveRolesFromUser(userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Profile").Preload("Roles").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Profile").Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Preload("Profile").Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) AssignRoleToUser(userRole *models.UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *userRepository) GetUserWithRoles(userID uint, user *models.User) error {
	return r.db.Preload("Roles").First(user, userID).Error
}

func (r *userRepository) RemoveRolesFromUser(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.UserRole{}).Error
}
