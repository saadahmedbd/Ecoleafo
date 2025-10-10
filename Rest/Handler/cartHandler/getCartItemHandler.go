package carthandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	cart, err := h.service.GetCart(buyerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, cart, 200)
}
