package db

import (
	"context"
	"database/sql"
	"time"
)

// New initializes a new database connection with the provided parameters.
func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)

	duration, err := time.ParseDuration(maxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(duration)
	db.SetMaxIdleConns(maxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
