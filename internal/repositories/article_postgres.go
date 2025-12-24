package repositories

import (
	"errors"
	"fmt"

	"github.com/MerBerd/blog-app/internal/models"
	"github.com/jmoiron/sqlx"
)

type ArticlePostgres struct {
	db *sqlx.DB
}

func NewArticlePostgres(db *sqlx.DB) *ArticlePostgres {
	return &ArticlePostgres{db: db}
}

func (r *ArticlePostgres) Create(userId int, article models.Article) (int, error) {

	var exists bool
	var articleID int

	query := fmt.Sprintf(" SELECT EXISTS (SELECT 1 FROM %s WHERE id = $1)", usersTable)
	err := r.db.Get(&exists, query, userId)
	if err != nil {
		return 0, err
	}

	if exists {
		createArticleQuery := fmt.Sprintf("INSERT INTO %s (author_id, title, content) values ($1, $2, $3) RETURNING id", articlesTable)
		if err = r.db.Get(&articleID, createArticleQuery,
			userId,
			article.Title,
			article.Content,
		); err != nil {
			return 0, err
		}

		return articleID, nil

	} else {
		return 0, errors.New("Such user does not exist")
	}

}

func (r *ArticlePostgres) GetAll(userId int) ([]models.Article, error) {
	var articles []models.Article

	query := fmt.Sprintf("SELECT id, title, content, created_at FROM %s WHERE author_id=$1", articlesTable)

	err := r.db.Select(&articles, query, userId)

	return articles, err
}

func (r *ArticlePostgres) GetById(userId, id int) (models.Article, error) {
	var article models.Article

	query := fmt.Sprintf("SELECT id, title, content, created_at FROM %s WHERE id=$1 AND author_id=$2", articlesTable)

	err := r.db.Get(&article, query, id, userId)

	return article, err
}
