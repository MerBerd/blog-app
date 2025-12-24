package repositories

import (
	"github.com/MerBerd/blog-app/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Article interface {
	Create(userId int, list models.Article) (int, error)
	GetAll(userId int) ([]models.Article, error)
	GetById(userId, id int) (models.Article, error)
}

type Comment interface {
}

type Repository struct {
	Authorization
	Article
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Article:       NewArticlePostgres(db),
	}
}
