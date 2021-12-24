package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-microservices-demo/domain"
	"github.com/leslesnoa/go-microservices-demo/service"
)

func sanitiyCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

func Start() {

	sanitiyCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
