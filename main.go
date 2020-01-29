package main

import (
	"fmt"
	productController "lab-golang-api/controllers/products"
)

func main () {
	fmt.Println("Hello World!!")
	productController.Get("GGG")
}