package models

import(
	"gopkg.in/mgo.v2/bson"
	"golang-mvc-webapp/config"
	"golang-mvc-webapp/db"
)

type ProductModel struct {}

type ProductItem struct {
	Sku string `json:"sku"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

var (
	DB *db.Mongodb
	dbName = config.Getenv("APP_MONGO_DATABASE")
)

func GetProductModel() *ProductModel {
	return &ProductModel{}
}

func (c *ProductModel) Create(p ProductItem) error {
	session := db.GetMongodb().GetSession()
	defer session.Close()

	err := session.DB(dbName).C("products").Insert(p)
	return err
}

func (c *ProductModel) All() ([]ProductItem, error) {

	session := db.GetMongodb().GetSession()
	defer session.Close()

	var results []ProductItem
	err := session.DB(dbName).C("products").Find(bson.M{}).All(&results)
	return results, err
}