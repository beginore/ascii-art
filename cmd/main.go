package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)

	mux.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	log.Print("starting server on: http://localhost:2929/")
	err := http.ListenAndServe(":2929", mux)
	log.Fatal(err)
}
