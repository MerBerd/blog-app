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

func (s *CommentService) GetById(commentId int) (models.Comment, error) {
	return s.repo.GetById(commentId)
}

func (s *CommentService) Update(userId, commentId int, input models.UpdateCommentInput) error {
	return s.repo.Update(userId, commentId, input)
}

func (s *CommentService) Delete(userId, commentId int) error {
	return s.repo.Delete(userId, commentId)
}
