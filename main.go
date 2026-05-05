package main

import (
	handler "goth-countries/handlers"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /", handler.CountryList)
	mux.HandleFunc("GET /Country/{name}", handler.Country)
	mux.HandleFunc("GET /search", handler.Search)

	// server stuff
	log.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
