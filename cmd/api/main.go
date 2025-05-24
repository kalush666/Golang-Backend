package main

import (
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

	// Initialize the storage layer
	store := store.NewStorage(nil)

	// Initialize the application with configuration
	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
		},
		store: store,
	}

	// Initialize the application and mount the routes
	mux := app.mount()
	log.Fatal(app.run(mux))
}
