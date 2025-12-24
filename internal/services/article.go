package services

import (
	"github.com/MerBerd/blog-app/internal/models"
	"github.com/MerBerd/blog-app/internal/repositories"
)

type ArticleService struct {
	repo repositories.Article
}

func NewArticleService(repo repositories.Article) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Create(userId int, list models.Article) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ArticleService) GetAll(userId int) ([]models.Article, error) {
	return s.repo.GetAll(userId)
}

func (s *ArticleService) GetById(userId, id int) (models.Article, error) {
	return s.repo.GetById(userId, id)
}
