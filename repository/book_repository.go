package repository

import (
	"context"

	"github.com/ssr0016/book/model"
)

type BookRepository interface {
	Save(ctx context.Context, book model.Book) error
	Update(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, bookID int64) error
	FindByID(ctx context.Context, bookID int64) (model.Book, error)
	FindAll(ctx context.Context) ([]model.Book, error)
}
