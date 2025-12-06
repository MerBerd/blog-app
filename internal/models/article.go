package models

type Article struct {
	Id        int    `json:"_"`
	AuthorId  int    `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type Attachment struct {
	Id         int    `json:"_"`
	ArticleId  int    `json:"article_id"`
	FilePath   string `json:"file_path"`
	UploadedAt string `json:"uploaded_at"`
}
