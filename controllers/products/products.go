package products

import(
	"fmt"
	"lab-golang-api/db"
	"lab-golang-api/models"
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