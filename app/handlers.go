package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-microservices-demo/service"
)

type Customer struct {
	Name   string `json:"full_name"`
	City   string `json:"city"`
	Zipcod string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Ashish", City: "New Delhi", Zipcod: "110075"},
	// 	{Name: "Rob", City: "New Delhi", Zipcod: "110075"},
	// }

	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// get path parameter
	vars := mux.Vars(r)
	// fmt.Fprint(w, vars["customer_id"])
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, err.Error())
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
