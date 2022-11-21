package app

import (
	"log"
	"net/http"
)

func Start() {
	// define routes
	http.HandleFunc("/", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}