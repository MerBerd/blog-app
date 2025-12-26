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
	Update(userId, articleId int, input models.UpdateArticleInput) error
	Delete(userId, articleId int) error
}

type Comment interface {
	GetAll(articleId int) ([]models.Comment, error)
	Create(userId, articleId int, input models.Comment) (int, error)
	GetById(commentId int) (models.Comment, error)
	Update(userId, commentId int, input models.UpdateCommentInput) error
	Delete(userId, commentId int) error
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
		Comment:       NewCommentsPostgres(db),
	}
}
