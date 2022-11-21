package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name string	`json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string`json:"zip_code" xml:"zipcode"`
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


func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
