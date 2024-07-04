package services

import (
	"errors"
	"final-project-golang-individu/models"
	"final-project-golang-individu/repositories"
	"final-project-golang-individu/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, password, email, role string) (*models.User, error)
	Login(username, password string) (string, error)
}

type authService struct {
	userRepository repositories.UserRepository
	roleRepository repositories.RoleRepository
}

func NewAuthService(userRepo repositories.UserRepository, roleRepo repositories.RoleRepository) AuthService {
	return &authService{
		userRepository: userRepo,
		roleRepository: roleRepo,
	}
}

func (s *authService) Register(username, password, email, role string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userRole, err := s.roleRepository.GetRoleByName(role)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Roles:    []models.Role{*userRole},
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	token, err := utils.GenerateToken(user.ID, roleNames)
	if err != nil {
		return "", err
	}

	return token, nil
}
