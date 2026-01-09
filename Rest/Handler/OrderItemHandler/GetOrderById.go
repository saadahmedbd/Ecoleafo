package orderitemhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderItemId := r.PathValue("orderitemId")
	sId, err := strconv.Atoi(orderItemId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var orderItems models.OrderItem
	result := Config.DB.Preload("Order").Preload("Product").Preload("Seller").First(&orderItems, sId)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}

	util.SendData(w, orderItems, 200)
}
