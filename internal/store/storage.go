package store

import (
	"context"
	"database/sql"
)

// Storage struct holds the storage interfaces for different entities in the application.
type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
	}
	Users interface {
		Create(ctx context.Context, user *User) error
	}
}

// newPostgresStorage initializes a new Storage instance with Postgres implementations for the storage interfaces.
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
