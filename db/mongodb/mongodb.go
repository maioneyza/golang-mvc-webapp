package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"

)

type DataStore struct {
	Session *mgo.Session
}

var (
	Session *mgo.Session
	err error
)

func init() {
	fmt.Println("mongodb#init")

	Session, err = mgo.Dial("mongodb://mongo")
	Session.SetMode(mgo.Monotonic, true)
	
	if err != nil {
		fmt.Println("ERR mgo.Dial:", err)
		panic(err)
	}

	fmt.Println("mongodb#init: success")
}

func Ds() *DataStore {
	return &DataStore{Session.Copy()}
}

