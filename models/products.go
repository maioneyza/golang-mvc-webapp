package models

import(
	"gopkg.in/mgo.v2"
	"golang-mvc-webapp/config"
)

type Products struct {
	Connection *mgo.Session
}

type ProductItem struct {
	Sku string `json:"sku"`
	Name string `json:"name":`
	Price float32 `json:"price"`
}

func (c *Products) Create(p ProductItem) error {

	db := c.Connection.DB(config.Getenv("APP_MONGO_DATABASE"))
	productsCollection := db.C("products")
	err := productsCollection.Insert(p)
	return err
}

func (c *Products) All() ([]ProductItem, error) {
	var results []ProductItem
	db := c.Connection.DB(config.Getenv("APP_MONGO_DATABASE"))
	err := db.C("client").Find(nil).All(&results)
	return results, err
}

func (c *Products) CloseSession() {
	c.Connection.Close()
}