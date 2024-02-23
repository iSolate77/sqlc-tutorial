package core

import (
	"context"

	"github.com/jackc/pgx"
)

type QueryRepo interface {
	WithTx(tx pgx.Tx) *db.Queries
	CreateAuthor(ctx context.Context, author Author) (db.Author, error)
}
