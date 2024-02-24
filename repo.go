package main

import (
	"context"
	"log"

	"github.com/iSolate77/sqlc-tutorial/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Create(ctx context.Context, author Author) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	arg := db.CreateAuthorParams{
		Name: author.Name,
		Bio:  pgtype.Text{String: author.Bio, Valid: true},
	}

	log.Printf("arguments: %+v\n", arg)

	queries := db.New(tx)

	_, err = queries.CreateAuthor(ctx, arg)
	if err != nil {
		return err
	}

	log.Printf("repo: Creating author: %+v\n", author)

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
