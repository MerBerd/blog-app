package models

type Comment struct {
	Id        int    `json:"_"`
	ArticleId int    `json:"article_id"`
	AuthorId  int    `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
