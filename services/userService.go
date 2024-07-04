package services

import (
	"final-project-golang-individu/models"
	"final-project-golang-individu/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repository.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repository.GetUserByID(id)
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	return s.repository.GetUserByUsername(username)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repository.GetAllUsers()
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repository.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repository.DeleteUser(id)
}
