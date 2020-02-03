package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"golang-mvc-webapp/config"
	"log"
)

type Mongodb struct {}

var (
	session *mgo.Session
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
	
	session, err = mgo.Dial(dsn)

	if (err != nil) {
		log.Fatal("Cannot Connect Mongodb")
	}
}

func (c Mongodb) New() *Mongodb {
	return c
}

func (c *Mongodb) connection() *mgo.Session {
	 s := session.Copy()
	 return s
}