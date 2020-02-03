package users

import (
	"github.com/gorilla/mux"
)

func BindRoutes(r *mux.Router) {
	sr := r.PathPrefix("/users").Subrouter()
	sr.HandleFunc("", userIndexAction)
}