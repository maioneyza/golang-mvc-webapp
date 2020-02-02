package main

import (
	"github.com/gorilla/mux"
	"net/http"
	ProductsController "golang-mvc-webapp/controllers/products"
	UsersController "golang-mvc-webapp/controllers/users"
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