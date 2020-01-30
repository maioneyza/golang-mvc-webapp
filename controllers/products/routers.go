package products

import (
	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
	sr = r.PathPrefix("/products").Subrouter()
)

func BindRoutes() {
	sr.HandleFunc("", ListAction)
}