package database

import (
	"context"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const defaultTimeout = 3 * time.Second

type DB struct {
	*pgx.Conn
}

func New(dsn string, automigrate bool) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	db, err := pgx.Connect(ctx, "postgres://"+dsn)
	if err != nil {
		return nil, err
	}

	if automigrate {
		migrator, err := migrate.New("file://assets/migrations", "postgres://"+dsn)
		if err != nil {
			log.Fatal(err)
		}

		if err := migrator.Up(); err != nil && err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	return &DB{db}, nil
}
