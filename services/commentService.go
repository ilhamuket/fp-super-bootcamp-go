package services

import (
	"final-project-golang-individu/models"
	"final-project-golang-individu/repositories"
)

type CommentService interface {
	CreateComment(comment *models.Comment) error
	GetCommentByID(id uint) (*models.Comment, error)
	GetCommentsByNewsID(newsID uint) ([]models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id uint) error
}

type commentService struct {
	repository repositories.CommentRepository
}

func NewCommentService(repository repositories.CommentRepository) CommentService {
	return &commentService{repository: repository}
}

func (s *commentService) CreateComment(comment *models.Comment) error {
	return s.repository.CreateComment(comment)
}

func (s *commentService) GetCommentByID(id uint) (*models.Comment, error) {
	return s.repository.GetCommentByID(id)
}

func (s *commentService) GetCommentsByNewsID(newsID uint) ([]models.Comment, error) {
	return s.repository.GetCommentsByNewsID(newsID)
}

func (s *commentService) UpdateComment(comment *models.Comment) error {
	return s.repository.UpdateComment(comment)
}

func (s *commentService) DeleteComment(id uint) error {
	return s.repository.DeleteComment(id)
}
