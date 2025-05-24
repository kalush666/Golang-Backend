package store

import (
	"context"
	"database/sql"
)

// Storage struct holds the storage interfaces for different entities in the application.
type PostStore struct {
	db *sql.DB
}

// NewPostStore initializes a new PostStore with the given database connection.
func (s *PostStore) Create(ctx context.Context) error {
	return nil
}
