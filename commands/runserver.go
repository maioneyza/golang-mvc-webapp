package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/Session"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	Name string
	Phone string
}

type Users []User

func main() {
	r := mux.NewRouter()

	session, err := getConnection();

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")
	db.Login("anuchit","$$jak$$")
	c := db.C("phonebook")
	err = c.Insert(&User{"Ale", "+55 53 8116 9639"}, &User{"Cla", "+55 53 8402 8510"})
	
	if err != nil {
			log.Fatal(err)
	}

	result := User{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("hello")
		fmt.Fprintf(w, "Hello Worddddld!!!")
	})

	r.HandleFunc("/test", testHandle)

	fmt.Println("Server listening!")
	http.ListenAndServe(":80", r)
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	// Query All
	var results []Person
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
    json.NewEncoder(w).Encode(todos)
}

function getConnection() (*Session, error){
	session, err := mgo.Dial("mongodb://mongo:27017/local")
	defer session.Close()
	return session, err
}