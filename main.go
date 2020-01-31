package main

import (
	"github.com/gorilla/mux"
	"net/http"
	ProductsController "lab-golang-api/controllers/products"
	UsersController "lab-golang-api/controllers/users"
	"log"
)

func main () {
	r := mux.NewRouter()

	ProductsController.BindRoutes(r)
	UsersController.BindRoutes(r)
	
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal("Serving error.", err)
	}
}