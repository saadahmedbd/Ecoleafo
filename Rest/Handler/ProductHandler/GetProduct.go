package producthandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid Method", http.StatusBadRequest)
		return
	}
	var Products []models.Product

	err := Config.DB.Preload("Seller").Preload("Category").Preload("CartItems").Preload("OrderItems").Preload("Reviews").Find(&Products).Error
	if err != nil {
		http.Error(w, "Failed fetch user", http.StatusInternalServerError)
		return
	}
	util.SendData(w, Products, 200)
}
