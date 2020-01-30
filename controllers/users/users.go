package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type User struct{
	ID string
	Name string
}

var (
	Users []User
)

func RegisterRoutes(r *mux.Subrouter) {
	r.handleFunc("/", indexAction).Methods("GET")
}

func indexAction(w http.ResponseWriter, r *http.Request) {
	users = []Users{
		&User{"1", "Anuchit Thiamuan"},
		&User{"2", "Jak Jak"}
	}
	
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(users)
}

