package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ssr0016/book/data/response"
	"github.com/ssr0016/book/helper"
	"github.com/ssr0016/book/model"
)

type BookRepositoryImple struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImple{Db: Db}
}

func (n *BookRepositoryImple) Save(ctx context.Context, book model.Book) error {
	tx, err := n.Db.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		INSERT INTO books
		(title, description, author, published_at)
		VALUES
		($1, $2, $3, $4)
	`

	_, err = tx.ExecContext(ctx, rawSQL, book.Title, book.Description, book.Author, book.PublishedAt)
	if err != nil {
		return err
	}

	return nil
}

func (n *BookRepositoryImple) FindByID(ctx context.Context, bookID int64) (model.Book, error) {
	tx, err := n.Db.Begin()
	if err != nil {
		return model.Book{}, err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		SELECT
			id, title, description, author, published_at
		FROM
			books
		WHERE
			id = $1
	`
	result, err := tx.QueryContext(ctx, rawSQL, bookID)
	if err != nil {
		return model.Book{}, err
	}
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.PublishedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return book, response.ErrBookNotFound
			}
		} else {
			return book, nil
		}
	}

	return book, nil
}

func (n *BookRepositoryImple) FindAll(ctx context.Context) ([]model.Book, error) {
	tx, err := n.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		SELECT
			id, title, description, author, published_at
		FROM
			books
	`
	result, err := tx.QueryContext(ctx, rawSQL)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.PublishedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (n *BookRepositoryImple) Update(ctx context.Context, book model.Book) error {
	tx, err := n.Db.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		UPDATE
			books
		SET
			title = $1, description = $2, author = $3, published_at = $4
		WHERE
			id = $5
	`

	_, err = tx.ExecContext(ctx, rawSQL, book.Title, book.Description, book.Author, book.PublishedAt, book.ID)
	if err != nil {
		return err
	}

	return nil
}

func (n *BookRepositoryImple) Delete(ctx context.Context, bookID int64) error {
	tx, err := n.Db.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		DELETE FROM
			books
		WHERE
			id = $1
	`
	_, err = tx.ExecContext(ctx, rawSQL, bookID)
	if err != nil {
		return err
	}

	return nil
}
