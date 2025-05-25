package main

import (
	"BackendLearning/internal/db"
	"BackendLearning/internal/env"
	"BackendLearning/internal/store"
	"github.com/joho/godotenv"
	"log"
)

// this go file will serve as setting up configuration, initializing the application, and starting the server
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	//set up the database configuration
	cfg := dbConfig{
		addr:         env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/social?sslmode=disable"),
		maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"),
	}

	// Initialize the database connection
	db, err := db.New(cfg.addr, cfg.maxOpenConns, cfg.maxIdleConns, cfg.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

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
