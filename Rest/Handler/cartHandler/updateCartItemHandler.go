package carthandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	productIDStr := r.PathValue("productId")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var req cartitem.UpdateCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCartItem(buyerID, uint(productID), req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.SendData(w, map[string]string{"message": "Cart item updated"}, 200)
}
