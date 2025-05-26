package main

import (
	"BackendLearning/internal/db"
	"BackendLearning/internal/env"
	"BackendLearning/internal/store"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

// this go file will serve as setting up configuration, initializing the application, and starting the server
func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	dbUser := env.GetString("DB_USER", "postgres")
	dbHost := env.GetString("DB_HOST", "localhost")
	dbName := env.GetString("DB_NAME", "social")

	cfg := dbConfig{
		addr: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbUser,
			env.GetString("DB_PASSWORD", ""),
			dbHost,
			env.GetString("DB_PORT", "5432"),
			dbName,
		),
		maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	}

	// Initialize the database connection
	db, err := db.New(cfg.addr, cfg.maxOpenConns, cfg.maxIdleConns, cfg.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("db coneccted")

	// Initialize the storage layer
	store := store.NewStorage(db)

	// Initialize the application with configuration
	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
			db:   cfg,
		},
		store: store,
	}

	// Initialize the application and mount the routes
	mux := app.mount()
	log.Fatal(app.run(mux))
}
