package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

var (
	MongodbConnection *mgo.Session
	err error
)

func init() {
	fmt.Println("Connecting DB....")
	MongodbConnection, err = mgo.Dial("mongodb://mongo")

	if (err != nil) {
		log.Fatal("Cannot Connect Mongodb")
	}
}

func GetMongodbSession() *mgo.Session {
	return MongodbConnection.Copy()
}