package models

import "time"

type Comment struct {
	ID        int       `db:"id" json:"id"`
	Content   string    `db:"content" json:"content"`
	AuthorID  int       `db:"author_id" json:"-"`
	ArticleID int       `db:"article_id" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
