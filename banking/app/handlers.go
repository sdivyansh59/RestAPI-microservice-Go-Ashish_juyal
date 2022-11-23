package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/sdivyansh59/banking/service"
)

type Customer struct {
	Name string	`json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string`json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch* CustomerHandlers)getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// this is hard coded value
	// customers := [] Customer {
	// 	{"Ashis", "New Delhi", "100247"},
	// 	{"Rob", "New Delhi", "273201"},
	// }

	// ask data from service
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

