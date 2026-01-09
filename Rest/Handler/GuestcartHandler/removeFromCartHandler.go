package guestcarthandler

import (
	"fmt"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// DELETE /api/cart/items/{product_id}
func (h *GuestcartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	sessionID := h.getOrCreateSessionID(r, w)
	userID := h.getUserID(r)

	productId := r.PathValue("product_id")
	productID, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "cannot convert id to int", http.StatusBadRequest)
		return
	}

	if err := h.service.RemoveCartItem(sessionID, userID, uint(productID)); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Item removed from cart",
	}
	util.SendData(w, response, http.StatusOK)
}
