package repositories

import (
	"final-project-golang-individu/models"
	"github.com/jinzhu/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetCommentByID(id uint) (*models.Comment, error)
	GetCommentsByNewsID(newsID uint) ([]models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) GetCommentsByNewsID(newsID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("news_id = ?", newsID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) UpdateComment(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) DeleteComment(id uint) error {
	return r.db.Delete(&models.Comment{}, id).Error
}
