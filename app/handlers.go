package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name   string `json:"full_name"`
	City   string `json:"city"`
	Zipcod string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ashish", City: "New Delhi", Zipcod: "110075"},
		{Name: "Rob", City: "New Delhi", Zipcod: "110075"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
