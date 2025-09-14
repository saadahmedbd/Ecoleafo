package orderhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please provide valid method", http.StatusBadRequest)
		return
	}
	UpdateOrderId := r.PathValue("orderId")
	id, err := strconv.Atoi(UpdateOrderId)

	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}
	var existingOrder models.Order
	if err := Config.DB.First(&existingOrder, id).Error; err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}
	var UpdateOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&UpdateOrder); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Update existing seller with new values
	existingOrder.OrderNumber = UpdateOrder.OrderNumber
	existingOrder.Status = UpdateOrder.Status
	existingOrder.Total = UpdateOrder.Total
	existingOrder.PaymentStatus = UpdateOrder.PaymentStatus
	existingOrder.ShippingAddress = UpdateOrder.ShippingAddress
	existingOrder.CustomerEmail = UpdateOrder.CustomerEmail
	existingOrder.CustomerPhone = UpdateOrder.CustomerPhone
	existingOrder.Notes = UpdateOrder.Notes

	if err := Config.DB.Save(&existingOrder).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}
	util.SendData(w, UpdateOrder, 200)
}
