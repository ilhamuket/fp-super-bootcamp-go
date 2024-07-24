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
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) CommentService {
	return &commentService{commentRepo: commentRepo}
}

func (s *commentService) CreateComment(comment *models.Comment) error {
	return s.commentRepo.CreateComment(comment)
}

func (s *commentService) GetCommentByID(id uint) (*models.Comment, error) {
	return s.commentRepo.GetCommentByID(id)
}

func (s *commentService) GetCommentsByNewsID(newsID uint) ([]models.Comment, error) {
	return s.commentRepo.GetCommentsByNewsID(newsID)
}

func (s *commentService) UpdateComment(comment *models.Comment) error {
	return s.commentRepo.UpdateComment(comment)
}

func (s *commentService) DeleteComment(id uint) error {
	return s.commentRepo.DeleteComment(id)
}
