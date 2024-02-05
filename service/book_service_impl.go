package service

import (
	"context"
	"time"

	"github.com/ssr0016/book/data/request"
	"github.com/ssr0016/book/data/response"
	"github.com/ssr0016/book/model"
	"github.com/ssr0016/book/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookRepositoryImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (n *BookServiceImpl) Create(ctx context.Context, req request.BookCreateRequest) error {

	result, err := n.BookRepository.BookTaken(ctx, 0, req.Title)
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return response.ErrBookNameAlreadyTaken
	}

	err = n.BookRepository.Save(ctx, model.Book{
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		PublishedAt: time.Now().UTC().Format(time.RFC3339Nano),
	})
	if err != nil {
		return err
	}

	return nil
}

func (n *BookServiceImpl) GetByID(ctx context.Context, bookID int64) response.BookResponse {
	book, err := n.BookRepository.FindByID(ctx, bookID)
	if err != nil {
		panic(err)
	}

	return response.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		PublishedAt: book.PublishedAt,
	}
}

func (n *BookServiceImpl) Search(ctx context.Context) ([]response.BookResponse, error) {

	books, err := n.BookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var bookResp []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{
			ID:          value.ID,
			Title:       value.Title,
			Description: value.Description,
			Author:      value.Author,
			PublishedAt: value.PublishedAt,
		}
		bookResp = append(bookResp, book)
	}

	return bookResp, nil
}

func (n *BookServiceImpl) Update(ctx context.Context, req request.BookUpdateRequest) error {

	result, err := n.BookRepository.BookTaken(ctx, req.ID, req.Title)
	if err != nil {
		return err
	}

	if len(result) > 1 || (len(result) == 1 && result[0].ID != req.ID) {
		return response.ErrBookNotFound
	}

	err = n.BookRepository.Update(ctx, model.Book{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		PublishedAt: time.Now().UTC().Format(time.RFC3339Nano),
	})
	if err != nil {
		return err
	}

	return nil
}

func (n *BookServiceImpl) Delete(ctx context.Context, bookID int64) error {
	book, err := n.BookRepository.FindByID(ctx, bookID)
	if err != nil {
		return err
	}

	if book.ID == 0 {
		return response.ErrBookNotFound
	}

	err = n.BookRepository.Delete(ctx, bookID)
	if err != nil {
		return err
	}

	return nil
}
