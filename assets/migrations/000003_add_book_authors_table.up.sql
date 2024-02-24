CREATE TABLE book_authors (
    book_id INT,
    author_id INT,
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (author_id) REFERENCES authors(id),
    PRIMARY KEY (book_id, author_id)
);
