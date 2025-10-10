package carthandler

import (
	"encoding/json"
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) AddToWishlist(w http.ResponseWriter, r *http.Request) {
	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	var req struct {
		ProductID uint `json:"product_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.AddToWishlist(buyerID, req.ProductID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.SendData(w, map[string]string{"message": "Item added to wishlist"}, 200)
}
