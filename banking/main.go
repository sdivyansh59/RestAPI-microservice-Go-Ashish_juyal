package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string	`json:"full_name"`
	City string `json:"city"`
	Zipcode string`json:"zip_code"`
}


func main() {
	// define routes
	http.HandleFunc("/", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000",nil))
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := [] Customer {
		{"Ashis", "New Delhi", "100247"},
		{"Rob", "New Delhi", "273201"},
	}
	w.Header().Add("Content-Type", "application/json")
	// we have list of customers So let's Encode them in JSON format
	json.NewEncoder(w).Encode(customers)
}