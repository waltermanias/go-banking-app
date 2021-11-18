package app

import (
	"banking-app/domain"
	"banking-app/errs"
	"banking-app/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	var customers []domain.Customer
	var err *errs.AppError

	if len(status) == 0 {
		customers, err = ch.service.GetAllCustomer("")
	} else {
		customers, err = ch.service.GetAllCustomer(status)
	}

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) getById(w http.ResponseWriter, r *http.Request) {

	customer, err := ch.service.GetById(mux.Vars(r)["customer_id"])

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func createCustomer(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Post request received")
}
