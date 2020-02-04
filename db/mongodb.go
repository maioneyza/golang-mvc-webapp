package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"golang-mvc-webapp/config"
	"log"
)

type Mongodb struct {
	session *mgo.Session
}

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

func GetMongodb() *Mongodb {
	return &Mongodb{session.Copy()}
}

func (c *Mongodb) GetSession() *mgo.Session {
	return c.session
}

func (c *Mongodb) CloseSession() {
	c.session.Close()
}