package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	// define routes
	http.HandleFunc("/", handler)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000",nil))
}