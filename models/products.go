package models

import(
	"gopkg.in/mgo.v2"
	"log"
)

type Products struct {
	Connection *mgo.Session
}

type ProductItem struct {
	Sku string
	Name string
	Price float32
}

// type ProductItems []ProductItem

func (c *Products) Create() bool {
	db := c.Connection.DB("test")
	t := db.Login("anuchit", "$$jak$$")

	log.Println(t)
	
	productsCollection := db.C("products")
	
	err := productsCollection.Insert(
		&ProductItem{"123", "สบู่่ถ่านไม้ไผ่ใข้ขัดตัว 1", 29.53},
		&ProductItem{"124", "สบู่่ถ่านไม้ไผ่ใข้ขัดตัว 2", 20.25},
		&ProductItem{"125", "สบู่่ถ่านไม้ไผ่ใข้ขัดตัว 3", 27.55},
	)

	if (err != nil) {
		log.Fatal(err)
	}

	return (err == nil)
}

func (c *Products) CloseSession() {
	c.Connection.Close()
}