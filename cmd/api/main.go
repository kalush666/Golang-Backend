package main

import "log"

// this go file will serve as setting up configuration, initializing the application, and starting the server
func main() {
	app := &application{
		config: config{
			addr: ":8080",
		},
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
