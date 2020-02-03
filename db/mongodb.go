package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"golang-mvc-webapp/config"
	"log"
)

var (
	MongodbConnection *mgo.Session
	err error
)

func init() {
	fmt.Println("Connecting DB....")

	dsn := fmt.Sprintf(
		"mongodb://%s:%s@mongo/%s", 
		config.Getenv("APP_MONGO_USERNAME"), 
		config.Getenv("APP_MONGO_PASSWORD"), 
		config.Getenv("APP_MONGO_DATABASE"),
	)

	fmt.Println(dsn)

	MongodbConnection, err = mgo.Dial(dsn)

	if (err != nil) {
		log.Fatal("Cannot Connect Mongodb")
	}
}

func GetMongodbSession() *mgo.Session {
	return MongodbConnection.Copy()
}

