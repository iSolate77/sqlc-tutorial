package core

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookAuthor struct {
	BookID   int `json:"book_id"`
	AuthorID int `json:"author_id"`
}

type BookAuthorResponse struct {
	BookID   int    `json:"book_id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	AuthorID int    `json:"author_id"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}

type BookAuthorRequest struct {
	BookID   int `json:"book_id"`
	AuthorID int `json:"author_id"`
}
