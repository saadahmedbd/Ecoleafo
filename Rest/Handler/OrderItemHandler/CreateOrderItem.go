package orderitemhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	var orderItems models.OrderItem

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&orderItems)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	if orderItems.Quantity <= 0 {
		http.Error(w, "Quantity must be greater than 0", http.StatusBadRequest)
		return
	}
	result := Config.DB.Create(&orderItems)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, orderItems, 200)
}
