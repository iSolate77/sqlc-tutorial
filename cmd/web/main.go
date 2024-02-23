package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"sqlc-tutorial/cmd/api"
	"sqlc-tutorial/core"
	"sqlc-tutorial/internal/database"
	"sqlc-tutorial/internal/env"
	"sqlc-tutorial/internal/version"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn         string
		automigrate bool
	}
}

type application struct {
	config config
	db     *pgxpool.Pool
	logger *slog.Logger
	wg     sync.WaitGroup
	api    *api.AuthorController
}

func run(logger *slog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = env.GetString("DB_DSN", "mfaris@localhost:5432/sqlc?sslmode=disable")
	cfg.db.automigrate = env.GetBool("DB_AUTOMIGRATE", true)

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	dbPool, err := pgxpool.New(context.Background(), cfg.db.dsn)
	if err != nil {
		return err
	}
	fmt.Println("Connected to database")
	defer dbPool.Close()

	queries := db.New(dbPool)
	repo := core.NewRepo(queries, dbPool)
	authorService := core.NewService(repo)

	app := &application{
		config: cfg,
		db:     dbPool,
		logger: logger,
		api:    nil,
	}
	app.api = api.NewAuthorController(authorService)

	return app.serveHTTP()
}
