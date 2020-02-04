package products

import (
	"github.com/gorilla/mux"
)

func BindRoutes(r *mux.Router)  {
	r.HandleFunc("/products", indexAction).Methods("GET")
	
	sr := r.PathPrefix("/products").Subrouter()
	sr.HandleFunc("", createAction).Methods("POST")
}