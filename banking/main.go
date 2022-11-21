package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string	`json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string`json:"zip_code" xml:"zipcode"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}