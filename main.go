package main

import (
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	ProductsController "golang-mvc-webapp/controllers/products"
	UsersController "golang-mvc-webapp/controllers/users"
)

func main () {
	loadEnv()
	r := mux.NewRouter()

	ProductsController.BindRoutes(r)
	UsersController.BindRoutes(r)
	
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal("Serving error.", err)
	}
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}