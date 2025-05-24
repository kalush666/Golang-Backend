package main

import (
	"BackendLearning/internal/env"
	"log"
)

// this go file will serve as setting up configuration, initializing the application, and starting the server
func main() {
	// Initialize the application with configuration
	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
		},
	}

	// Initialize the application and mount the routes
	mux := app.mount()
	log.Fatal(app.run(mux))
}
