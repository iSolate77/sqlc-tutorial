package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbUrl := "postgres://mfaris@localhost:5432/sqlc?sslmode=disable"
	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}
	defer db.Close()

	repo := NewRepo(db)
	authorService := NewService(repo)
	authorController := newController(*authorService)
	http.HandleFunc("/authors", authorController.CreateAuthor)

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}
	log.Println("Starting server on port 8080...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on port 8080: %v\n", err)
	}
}
