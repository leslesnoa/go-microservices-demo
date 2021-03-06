package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-microservices-demo/dto"
	"github.com/leslesnoa/go-microservices-demo/service"
)

type AccountHandler struct {
	service service.AccoutService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
