// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Author struct {
	ID   int32
	Name string
	Bio  pgtype.Text
}

type Book struct {
	ID     int32
	Title  string
	Author pgtype.Text
	Year   pgtype.Int4
}

type BookAuthor struct {
	BookID   int32
	AuthorID int32
}
