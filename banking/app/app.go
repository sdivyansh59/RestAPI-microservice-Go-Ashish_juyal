package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// creating our own multiplexer
	// mux := http.NewServeMux()

	// gorilla mux router : 3rd party router  
	router := mux.NewRouter()
	// define routes
	router.HandleFunc("/", greet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)
    // it will take only int as id

	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Post request received")
}

