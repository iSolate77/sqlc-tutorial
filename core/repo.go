package core

import (
	"context"
	"fmt"

	"sqlc-tutorial/core/db"

	"github.com/jackc/pgx/v5"
)

type QueryRepo interface {
	WithTx(tx pgx.Tx) *db.Queries
	CreateAuthor(ctx context.Context, arg db.CreateAuthorParams) (db.Author, error)
	GetAuthor(ctx context.Context, id int) (db.Author, error)
	UpdateAuthor(ctx context.Context, arg db.Author) (db.Author, error)
	DeleteAuthor(ctx context.Context, id int) (db.Author, error)
	ListAuthors(ctx context.Context, limit int, offset int) ([]db.Author, error)
	CreateBook(ctx context.Context, arg db.Book) (db.Book, error)
	GetBook(ctx context.Context, id int) (db.Book, error)
	UpdateBook(ctx context.Context, arg db.Book) (db.Book, error)
	DeleteBook(ctx context.Context, id int) (db.Book, error)
	ListBooks(ctx context.Context, limit int, offset int) ([]db.Book, error)
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

func (r *Repo) Create(ctx context.Context, author Author) (Author, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return Author{}, fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	queries := r.ar.WithTx(tx)
	author, err = queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: author.Name,
		Bio:  author.Bio,
	})
	if err != nil {
		return Author{}, fmt.Errorf("could not create author: %w", err)
	}
	return author, nil
}
