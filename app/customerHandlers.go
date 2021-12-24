package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-microservices-demo/service"
)

// type Customer struct {
// 	Name   string `json:"full_name"`
// 	City   string `json:"city"`
// 	Zipcod string `json:"zip_code"`
// }
/*
*Try*
APIはJSON応答のみを返す必要があります
サーバーから予期しないエラーが発生した場合、APIはHTTPステータスコード500を返す必要があります
サーバーから顧客を正常に取得した場合、APIはHTTPステータスコード200を返す必要があります
*/

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// get path parameter
	vars := mux.Vars(r)
	// fmt.Fprint(w, vars["customer_id"])
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
