package store

import (
	"context"
	"database/sql"
)

// User struct represents a user in the system.
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

// UserStore struct holds the database connection for user-related operations.
type UserStore struct {
	db *sql.DB
}

// NewUserStore initializes a new UserStore with the given database connection.
func (s *UserStore) Create(ctx context.Context, user *User) error {

	query := `INSERT INTO users (username, email, password) VALUES($1, $2, $3) RETURNING id, created_at`

	

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
