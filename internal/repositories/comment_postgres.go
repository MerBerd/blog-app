package repositories

import (
	"errors"
	"fmt"

	"github.com/MerBerd/blog-app/internal/models"
	"github.com/jmoiron/sqlx"
)

type CommentsPostgres struct {
	db *sqlx.DB
}

func NewCommentsPostgres(db *sqlx.DB) *CommentsPostgres {
	return &CommentsPostgres{db: db}
}

func (r *CommentsPostgres) GetAll(articleId int) ([]models.Comment, error) {
	var comments []models.Comment

	query := fmt.Sprintf("SELECT id, content, created_at FROM %s WHERE article_id=$1", commentsTable)

	err := r.db.Select(&comments, query, articleId)

	return comments, err
}

func (r *CommentsPostgres) Create(userId, articleId int, input models.Comment) (int, error) {
	var exists bool
	var commentID int

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE id=$1)", articlesTable)
	err := r.db.Get(&exists, query, userId)
	if err != nil {
		return 0, err
	}

	if exists {
		createCommentQuery := fmt.Sprintf("INSERT INTO %s (article_id, author_id, content) values ($1, $2, $3) RETURNING id", commentsTable)
		if err = r.db.Get(&commentID, createCommentQuery,
			articleId,
			userId,
			input.Content,
		); err != nil {
			return 0, err
		}

		return commentID, nil

	} else {
		return 0, errors.New("Such article does not exist")
	}
}
