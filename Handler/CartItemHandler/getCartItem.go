package cartitemhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid Method", http.StatusBadRequest)
		return
	}
	var CartItems []models.CartItem

	err := Config.DB.Preload("Buyer").Preload("Product").Find(&CartItems).Error
	if err != nil {
		http.Error(w, "Failed fetch cartItem", http.StatusInternalServerError)
		return
	}
	util.SendData(w, CartItems, 200)

}
