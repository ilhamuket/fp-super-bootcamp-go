package services

import (
	"final-project-golang-individu/models"
	"final-project-golang-individu/repositories"
)

type NewsService interface {
	CreateNews(news *models.News) error
	GetNewsByID(id uint) (*models.News, error)
	GetAllNews() ([]models.News, error)
	UpdateNews(news *models.News) error
	DeleteNews(id uint) error
}

type newsService struct {
	repository repositories.NewsRepository
}

func NewNewsService(repository repositories.NewsRepository) NewsService {
	return &newsService{repository: repository}
}

func (s *newsService) CreateNews(news *models.News) error {
	return s.repository.CreateNews(news)
}

func (s *newsService) GetNewsByID(id uint) (*models.News, error) {
	return s.repository.GetNewsByID(id)
}

func (s *newsService) GetAllNews() ([]models.News, error) {
	return s.repository.GetAllNews()
}

func (s *newsService) UpdateNews(news *models.News) error {
	return s.repository.UpdateNews(news)
}

func (s *newsService) DeleteNews(id uint) error {
	return s.repository.DeleteNews(id)
}
