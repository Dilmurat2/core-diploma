package main

import (
	"core/src/app"
	"log"
	"net/http"
)

func main() {
	// init server
	server, err := app.InitServer()
	if err != nil {
		log.Fatal(err)
	}
	// start HTTP
	log.Printf("Starting server at port 8080")
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
}
