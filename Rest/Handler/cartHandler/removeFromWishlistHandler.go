package carthandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) RemoveFromWishlist(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.RemoveFromWishlist(buyerID, uint(productID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"message": "Item removed from wishlist"}, 200)
}
