package users

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func indexAction (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Index Page")
}

func listAction (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User List action")
}

func infoAction (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Fprintf(w, "User %s's info action", params["id"])
}