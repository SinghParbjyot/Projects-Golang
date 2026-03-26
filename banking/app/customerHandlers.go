package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/singhparbjyot/banking/service"
)

// Handler instance a service , a handler is in charge to recibe the http request extract the values and calls the service
type CustomerHandlers struct {
	service service.CustomerService
}

// Function which get the request from the client and gets the response to the server with the request that the client has remited
// and returns all the customers from the db in json or xml depends on the client in whcih format he wants.
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	//Handler calls the service
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		//this is for set or return  the content-type in format json
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		//Functions that encode to json our struct
		json.NewEncoder(w).Encode(customers)
	}

}

// Function returns the customer with the following id that client requested other response with the correspond error
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, errors := ch.service.GetCustomer(id)

	if errors != nil {
		// EXPLICACIÓN: Primero definimos que enviaremos JSON, luego el código de error.
		writeResponse(w, errors.Code, errors.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

// Method to optmize the repetead code  in the previous method
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
