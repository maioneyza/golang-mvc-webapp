package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	ProductsController "golang-mvc-webapp/controllers/products"
	UsersController "golang-mvc-webapp/controllers/users"
	"log"
	"net/http"
)

func main() {
	loadEnv()
	r := mux.NewRouter()

	ProductsController.BindRoutes(r)
	UsersController.BindRoutes(r)
	//fix port
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Serving error.", err)
	}
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
