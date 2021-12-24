package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-microservices-demo/domain"
	"github.com/leslesnoa/go-microservices-demo/service"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
