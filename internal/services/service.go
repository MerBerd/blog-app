package services

import (
	"github.com/MerBerd/blog-app/internal/models"
	"github.com/MerBerd/blog-app/internal/repositories"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Article interface {
	Create(userId int, list models.Article) (int, error)
	GetAll(userId int) ([]models.Article, error)
	GetById(userId, id int) (models.Article, error)
	Update(userId, id int, input models.UpdateArticleInput) error
	Delete(userId, articleId int) error
}

type Comment interface {
	GetAll(articleId int) ([]models.Comment, error)
	Create(userId, articleId int, input models.Comment) (int, error)
	GetById(commentId int) (models.Comment, error)
	Update(userId, commentId int, input models.UpdateCommentInput) error
	Delete(userId, commentId int) error
}

type Service struct {
	Authorization
	Article
	Comment
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Article:       NewArticleService(repos.Article),
		Comment:       NewCommentService(repos.Comment),
	}
}
