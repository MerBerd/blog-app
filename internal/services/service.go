package services

import "github.com/MerBerd/blog-app/internal/repositories"

type Authorization interface {
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
	return &Service{}
}
