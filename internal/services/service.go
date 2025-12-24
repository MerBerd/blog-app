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
}

type Comment interface {
}

type Service struct {
	Authorization
	Article
	Comment
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
