package models

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"golang-mvc-webapp/config"
	"golang-mvc-webapp/db"
)

type ProductModel struct {
	Connection *mgo.Session
}

type ProductItem struct {
	Sku string `json:"sku"`
	Name string `json:"name":`
	Price float32 `json:"price"`
}

func (class *ProductModel) Create(p ProductItem) error {

	db := class.database().C("products")
	productsCollection.Close()
	err := productsCollection.Insert(p)

	return err
}

func (c *ProductModel) All() ([]ProductItem, error) {
	var results []ProductItem
	err := c.database().C("products").Find(bson.M{}).All(&results)

	return results, err
}