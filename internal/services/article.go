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

func (s *ArticleService) Update(userId, articleId int, input models.UpdateArticleInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, articleId, input)
}

func (s *ArticleService) Delete(userId, articleId int) error {
	return s.repo.Delete(userId, articleId)
}
