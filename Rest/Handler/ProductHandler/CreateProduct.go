package producthandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
	util "github.com/saadahmedbd/Treestore/Util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Get seller ID from context
	sellerID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	var products models.Product
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&products)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	products.SellerID = sellerID
	result := Config.DB.Create(&products)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, products, 200)

}
