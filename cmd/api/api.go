package main

import (
	"log"
	"net/http"
	"time"
)

//this file is where we will define the API endpoints and their handlers and where the application will live

// application struct holds the configuration for the application
type application struct {
	config config
}

// config struct holds the configuration for the application
type config struct {
	// addr is the TCP address that the application will listen on
	addr string
}

// NewApplication initializes a new application with the given configuration
func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	// Define the API endpoints and their handlers
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

// main function initializes the application and starts the server
func (app *application) run(mux *http.ServeMux) error {
	srv := http.Server{
		// Set the address and handler for the server
		Addr: app.config.addr,
		// Handler is the HTTP request handler to use for the server.
		Handler: mux,
		// Configure timeouts for the server
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}
