package products

import(
	"golang-mvc-webapp/db"
	"golang-mvc-webapp/models"
	"encoding/json"
	"net/http"
)

func createAction(w http.ResponseWriter, r *http.Request) {
	productModel := productsModel()
	defer productModel.CloseSession()

	var item models.ProductItem
	json.NewDecoder(r.Body).Decode(&item)
	productModel.Create(item);

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(item)
}

func indexAction(w http.ResponseWriter, r *http.Request) {
	productModel := productsModel()
	defer productModel.CloseSession()

	var results []models.ProductItem
	results, err := productModel.All();

	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	if(err != nil) {
		json.NewEncoder(w).Encode(err)
	}
	
	json.NewEncoder(w).Encode(results)
}

func productsModel() *models.Products {
	return &models.Products{db.MongodbConnection.Copy()}
}