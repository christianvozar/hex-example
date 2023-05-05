package postgres

import (
	"context"
	"database/sql"

	"github.com/christianvozar/hex-example/internal/domain"
)

type Listener struct {
	db *sql.DB
}

func NewListener() (*Listener, error) {
	// Initialize the Postgres connection
	db, err := sql.Open("postgres", "your_connection_string")
	if err != nil {
		return nil, err
	}

	return &Listener{db: db}, nil
}

func (l *Listener) Listen(ctx context.Context) (<-chan domain.Event, error) {
	// Implement the logic to listen for changes in the Postgres database
	// ...
	return nil, nil
}
