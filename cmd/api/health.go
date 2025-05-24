package main

import "net/http"

// healthCheckHandler is a simple handler that responds with "OK" to indicate that the service is healthy.
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
