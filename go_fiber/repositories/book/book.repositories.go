package repositories

import (
	"context"
	"database/sql"
	"go_fiber/schema"
)

type BookRepositories interface {
	FindAll(ctx context.Context) ([]*schema.Book, error)
}

type bookRepositries struct {
	PostGreDatabase *sql.DB
}

func NewBookRepositories(postGreDatabase *sql.DB) BookRepositories {
	return &bookRepositries{
		PostGreDatabase: postGreDatabase,
	}
}

// FindAll implements UserRepositories.
func (r *bookRepositries) FindAll(ctx context.Context) ([]*schema.Book, error) {
	var books []*schema.Book
	rows, err := r.PostGreDatabase.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book schema.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Email)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}
