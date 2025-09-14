package orderitemhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid Method", http.StatusBadRequest)
		return
	}
	var orderItems []models.OrderItem

	err := Config.DB.Preload("Order").Preload("Product").Preload("Seller").Find(&orderItems).Error
	if err != nil {
		http.Error(w, "Failed fetch order item", http.StatusInternalServerError)
		return
	}
	util.SendData(w, orderItems, 200)

}
