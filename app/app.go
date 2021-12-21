package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
