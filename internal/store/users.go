package store

import (
	"context"
	"database/sql"
)

// UserStore struct holds the database connection for user-related operations.
type UserStore struct {
	db *sql.DB
}

// NewUserStore initializes a new UserStore with the given database connection.
func (s *UserStore) Create(ctx context.Context) error {
	return nil
}
