package main

import "net/http"

//this file is where we will define the API endpoints and their handlers and where the application will live

type application struct {
	config config
}

type config struct {
	// addr is the TCP address that the application will listen on
	addr string
}

// main function initializes the application and starts the server
func (app *application) run() error {
	// Create a new ServeMux and register the routes
	mux := http.NewServeMux()

	srv := http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	return srv.ListenAndServe()
}
