package core

import "github.com/jackc/pgx/v5/pgtype"

type Author struct {
	Id   int         `json:"id"`
	Name string      `json:"name"`
	Bio  pgtype.Text `json:"bio"`
}

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookAuthor struct {
	BookId   int `json:"book_id"`
	AuthorId int `json:"author_id"`
}

