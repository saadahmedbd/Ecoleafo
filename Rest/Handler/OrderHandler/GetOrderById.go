package orderhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")
	sId, err := strconv.Atoi(orderId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var orders models.Order
	result := Config.DB.Preload("Buyer").Preload("OrderItems").First(&orders, sId)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	// not return password

	util.SendData(w, orders, 200)
}
