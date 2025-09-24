package orderitemhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please provide valid method", http.StatusBadRequest)
		return
	}
	UpdateOrderItemId := r.PathValue("orderitemId")
	id, err := strconv.Atoi(UpdateOrderItemId)

	if err != nil {
		http.Error(w, "Invalid Order item ID", http.StatusBadRequest)
		return
	}
	var existingOrderItem models.OrderItem
	if err := Config.DB.First(&existingOrderItem, id).Error; err != nil {
		http.Error(w, "order item not found", http.StatusNotFound)
		return
	}
	var UpdateOrderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&UpdateOrderItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Update existing seller with new values
	existingOrderItem.OrderID = UpdateOrderItem.OrderID
	existingOrderItem.ProductID = UpdateOrderItem.ProductID
	existingOrderItem.SellerID = UpdateOrderItem.SellerID
	existingOrderItem.Quantity = UpdateOrderItem.Quantity
	existingOrderItem.Price = UpdateOrderItem.Price
	existingOrderItem.Total = UpdateOrderItem.Total
	existingOrderItem.Commission = UpdateOrderItem.Commission
	existingOrderItem.SellerEarning = UpdateOrderItem.SellerEarning

	if err := Config.DB.Save(&existingOrderItem).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}
	util.SendData(w, UpdateOrderItem, 200)
}
