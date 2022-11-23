package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdivyansh59/banking/domain"
	"github.com/sdivyansh59/banking/service"
)

func Start() {
	
	router := mux.NewRouter()
	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	
	// starting server
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
