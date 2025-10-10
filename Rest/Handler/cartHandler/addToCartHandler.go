package carthandler

import (
	"encoding/json"
	"net/http"

	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userRole := r.Header.Get("user_role")
	if userIDStr == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if userRole != "[buyer]" {
		http.Error(w, "Only buyers can add products to cart", http.StatusForbidden)
		return
	}

	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	var req cartitem.AddToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cart, err := h.service.AddToCart(buyerID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.SendData(w, cart, 200)
}
