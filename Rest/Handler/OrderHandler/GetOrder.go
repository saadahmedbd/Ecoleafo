package orderhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid Method", http.StatusBadRequest)
		return
	}
	var orders []models.Order

	err := Config.DB.Preload("Buyer").Preload("OrderItems").Find(&orders).Error
	if err != nil {
		http.Error(w, "Failed fetch order", http.StatusInternalServerError)
		return
	}
	util.SendData(w, orders, 200)

}
