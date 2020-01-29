package products

import (
	"fmt"
	"lab-golang-api/db/mongodb"
)
type Product struct {
	Sku string 
	Name string 
	Price int 
}

func Get(s string) {
	ds := mongodb.Ds()
	defer ds.Session.Close()

	db := ds.Session.DB("test");
	db.Login("anuchit", "$$jak$$")

	phonebookCollection := db.C("products")
	err := phonebookCollection.Insert(&Product{"123456", "Sa Pai Pay", 11})

	if (err != nil) {
		panic(err)
	} else {
		fmt.Println("Saved!!")
	}

	fmt.Println("You want some %v ?", err)
}