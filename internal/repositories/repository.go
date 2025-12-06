package repositories

type Authorization interface {
}

type Article interface {
}

type Comment interface {
}

type Repository struct {
	Authorization
	Article
	Comment
}

func NewRepository() *Repository {
	return &Repository{}
}
