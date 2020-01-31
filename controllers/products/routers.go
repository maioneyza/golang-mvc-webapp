package products

import (
	"github.com/gorilla/mux"
)

func BindRoutes(r *mux.Router)  {
	sr := r.PathPrefix("/products").Subrouter()
	sr.HandleFunc("/create", createAction)
}