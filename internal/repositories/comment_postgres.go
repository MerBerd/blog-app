package repositories

import (
	"errors"
	"fmt"
	"strings"

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

func (r *CommentsPostgres) GetById(commentId int) (models.Comment, error) {
	var comment models.Comment

	query := fmt.Sprintf("SELECT id, content, created_at FROM %s WHERE id=$1", commentsTable)

	err := r.db.Get(&comment, query, commentId)

	return comment, err
}

func (r *CommentsPostgres) Update(userId, commentId int, input models.UpdateCommentInput) error {
	var allowed bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE author_id=$1)", commentsTable)

	if err := r.db.Get(&allowed, query, userId); err != nil {
		return err
	}

	if allowed {
		setValues := make([]string, 0)
		args := make([]interface{}, 0)
		argId := 1

		if input.Content != nil {
			setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
			args = append(args, *input.Content)
			argId++
		}

		setQuery := strings.Join(setValues, ", ")

		query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%v", commentsTable, setQuery, argId)
		args = append(args, commentId)

		_, err := r.db.Exec(query, args...)
		return err

	} else {
		return errors.New("You are not allowed to change this comment")
	}

}

func (r *CommentsPostgres) Delete(userId, commentId int) error {
	var allowed bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE author_id=$1)", commentsTable)

	if err := r.db.Get(&allowed, query, userId); err != nil {
		return err
	}

	if allowed {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", commentsTable)
		_, err := r.db.Exec(query, commentId)
		return err
	} else {
		return errors.New("You are not allowed to delete this article")
	}
}
