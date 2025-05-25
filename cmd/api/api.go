package main

import (
	"BackendLearning/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

//this file is where we will define the API endpoints and their handlers and where the application will live

// application struct holds the configuration for the application
type application struct {
	config config
	store  store.Storage
}

// config struct holds the configuration for the application
type config struct {
	// addr is the TCP address that the application will listen on
	addr string
	db   dbConfig
}

// dbConfig struct holds the database configuration for the application
type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

// NewApplication initializes a new application with the given configuration
func (app *application) mount() http.Handler {
	// Create a new Chi router
	r := chi.NewRouter()

	// Middleware for logging
	r.Use(middleware.Logger)
	// Middleware for logging requests
	r.Use(middleware.Recoverer)
	// Middleware for generating and attaching a unique request ID to each request
	r.Use(middleware.RequestID)
	// Middleware for handling real IPs, useful if the app is behind a reverse proxy
	r.Use(middleware.RealIP)

	// Middleware for handling CORS
	r.Use(middleware.Timeout(60 * time.Second))

	// Middleware for setting headers
	r.Route("/v1", func(r chi.Router) {
		// Set common headers for all routes under /v1
		r.Get("/health", app.healthCheckHandler)
	})

	//posts

	//users

	//auth

	return r
}

// main function initializes the application and starts the server
func (app *application) run(mux http.Handler) error {
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
