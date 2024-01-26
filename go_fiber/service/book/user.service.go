package service

import (
	"context"
	repositories "go_fiber/repositories/book"
	"go_fiber/schema"
)

type IBookService interface {
	FindAll() ([]*schema.Book, error)
}

type bookService struct {
	Repositories repositories.BookRepositories
}

// FindAll implements IBookService.
func (b *bookService) FindAll() ([]*schema.Book, error) {
	ctx := context.TODO()
	return b.Repositories.FindAll(ctx)
}

func NewUserService(repositories repositories.BookRepositories) IBookService {
	return &bookService{
		Repositories: repositories,
	}
}
