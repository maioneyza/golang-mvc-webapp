package main

import (
	ProductsController "lab-golang-api/controllers/products"
)

func main() {
	ProductsController.Get("Banana")
}