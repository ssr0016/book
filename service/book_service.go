package service

import (
	"context"

	"github.com/ssr0016/book/data/request"
	"github.com/ssr0016/book/data/response"
)

type BookService interface {
	Create(ctx context.Context, req request.BookCreateRequest) error
	GetByID(ctx context.Context, bookID int64) response.BookResponse
	Search(ctx context.Context) ([]response.BookResponse, error)
	Update(ctx context.Context, req request.BookUpdateRequest) error
	Delete(ctx context.Context, bookID int64) error
}
