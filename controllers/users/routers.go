package users

import (
	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
	sr = r.PathPrefix("/users").Subrouter()
)

func BindRoutes() {
	sr.HandleFunc("", indexAction)
	sr.HandleFunc("/{id}/info", infoAction)
}