package cartitemhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please provide valid method", http.StatusBadRequest)
		return
	}
	cartItemId := r.PathValue("cartitemId")
	id, err := strconv.Atoi(cartItemId)
	if err != nil {
		http.Error(w, "Invalid cart Item ID", http.StatusBadRequest)
		return
	}
	var existingCartItem models.CartItem
	if err := Config.DB.First(&existingCartItem, id).Error; err != nil {
		http.Error(w, "CartItem not found", http.StatusNotFound)
		return
	}
	var UpdateCartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&UpdateCartItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Update existing seller with new values
	existingCartItem.ProductID = UpdateCartItem.ProductID
	existingCartItem.Quantity = UpdateCartItem.Quantity
	existingCartItem.Price = UpdateCartItem.Price

	if err := Config.DB.Save(&existingCartItem).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}
	util.SendData(w, UpdateCartItem, 200)
}
