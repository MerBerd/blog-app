package services

import (
	"github.com/MerBerd/blog-app/internal/models"
	"github.com/MerBerd/blog-app/internal/repositories"
)

type CommentService struct {
	repo repositories.Comment
}

func NewCommentService(repo repositories.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) GetAll(articleId int) ([]models.Comment, error) {
	return s.repo.GetAll(articleId)
}

func (s *CommentService) Create(userId, articleId int, input models.Comment) (int, error) {
	return s.repo.Create(userId, articleId, input)
}
