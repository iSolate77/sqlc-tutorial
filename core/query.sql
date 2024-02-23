-- name: GetAuthor :one
SELECT * FROM authors WHERE id = $1;

-- name: GetAuthors :many
SELECT * FROM authors;

-- name: CreateAuthor :one
INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors SET name = $1, bio = $2 WHERE id = $3 RETURNING *;

-- name: DeleteAuthor :one
DELETE FROM authors WHERE id = $1 RETURNING *;

-- name: GetBooksByAuthor :many
SELECT b.* FROM books b
JOIN book_authors ba ON b.id = ba.book_id
WHERE ba.author_id = $1;

-- name: GetBooksByAuthorAndTitle :many
SELECT b.* FROM books b
JOIN book_authors ba ON b.id = ba.book_id
WHERE ba.author_id = $1 AND b.title = $2;

-- name: GetBooksByAuthorAndTitleAndYear :many
SELECT b.* FROM books b
JOIN book_authors ba ON b.id = ba.book_id
WHERE ba.author_id = $1 AND b.title = $2 AND b.year = $3;

-- name: GetBooksByAuthorAndYear :many
SELECT b.* FROM books b
JOIN book_authors ba ON b.id = ba.book_id
WHERE ba.author_id = $1 AND b.year = $2;

-- name: GetBooksByTitle :many
SELECT * FROM books WHERE title = $1;

-- name: GetBookByTitleAndYear :one
SELECT * FROM books WHERE title = $1 AND year = $2;

-- name: AssociateBookWithAuthor :exec
INSERT INTO book_authors (book_id, author_id) VALUES ($1, $2);

-- name: UpdateBook :one
UPDATE books SET title = $1, year = $2 WHERE id = $3 RETURNING *;

-- name: DeleteBook :one
DELETE FROM books WHERE id = $1 RETURNING *;

-- name: CreateBook :one
INSERT INTO books (title, year) VALUES ($1, $2) RETURNING *;

-- name: GetBook :one
SELECT * FROM books WHERE id = $1;
