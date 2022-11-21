package app

import (
	"log"
	"net/http"
)

func Start() {
	// creating our own multiplexer
	mux := http.NewServeMux()
	// define routes
	mux.HandleFunc("/", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}