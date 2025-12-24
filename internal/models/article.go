package models

type Article struct {
	Id        int    `json:"id" db:"id"`
	AuthorId  int    `json:"-" db:"author_id"`
	Title     string `json:"title" db:"title" binding:"required"`
	Content   string `json:"content" db:"content" binding:"required"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type Attachment struct {
	Id         int    `json:"id"`
	ArticleId  int    `json:"article_id"`
	FilePath   string `json:"file_path"`
	UploadedAt string `json:"uploaded_at"`
}
