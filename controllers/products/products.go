package products

import(
	"fmt"
	"golang-mvc-webapp/db"
	"golang-mvc-webapp/models"
	"net/http"
)

func createAction(w http.ResponseWriter, r *http.Request) {
	ProductsModel := dataStore()
	ProductsModel.Create();
	fmt.Fprint(w, "Created!!!")
}

func dataStore() *models.Products {
	return &models.Products{db.MongodbConnection.Copy()}
}