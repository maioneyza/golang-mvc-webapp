package products

import (
	"fmt"
	"net/http"
)

func ListAction (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Product Index Page")
}