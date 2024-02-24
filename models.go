package main

type Author struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type Book struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookAuthor struct {
	BookID   int `json:"book_id"`
	AuthorID int `json:"author_id"`
}
