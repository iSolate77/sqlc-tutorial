package core

import (
	"context"
	"fmt"

	db "sqlc-tutorial/internal/database"

	"github.com/jackc/pgx/v5"
)

type QueryRepo interface {
	WithTx(tx pgx.Tx) *db.Queries
	CreateAuthor(ctx context.Context, arg db.CreateAuthorParams) (db.Author, error)
	GetAuthor(ctx context.Context, id int32) (db.Author, error)
	// UpdateAuthor(ctx context.Context, arg db.Author) (db.Author, error)
	// DeleteAuthor(ctx context.Context, id int32) (db.Author, error)
	// ListAuthors(ctx context.Context, limit int, offset int) ([]db.Author, error)
	// CreateBook(ctx context.Context, arg db.Book) (db.Book, error)
	// GetBook(ctx context.Context, id int) (db.Book, error)
	// UpdateBook(ctx context.Context, arg db.Book) (db.Book, error)
	// DeleteBook(ctx context.Context, id int) (db.Book, error)
	// ListBooks(ctx context.Context, limit int, offset int) ([]db.Book, error)
}

type dbInterface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type Repo struct {
	ar QueryRepo
	db dbInterface
}

func NewRepo(ar QueryRepo, db dbInterface) *Repo {
	return &Repo{ar: ar, db: db}
}

func (r *Repo) Create(ctx context.Context, auth Author) (Author, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return Author{}, fmt.Errorf("could not begin transaction: %w", err)
	}

	queries := r.ar.WithTx(tx)
	author, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: auth.Name,
		Bio:  auth.Bio,
	})
	if err != nil {
		return Author{}, fmt.Errorf("could not create author: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return Author{}, fmt.Errorf("could not commit transaction: %w", err)
	}
	res, err := r.ar.GetAuthor(ctx, author.ID)
	if err != nil {
		return Author{}, fmt.Errorf("could not get author: %w", err)
	}
	return Author{
		Id:   int(res.ID),
		Name: res.Name,
		Bio:  res.Bio,
	}, nil
}
