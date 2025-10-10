package carthandler

import (
	"encoding/json"
	"net/http"

	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) BulkUpdateCart(w http.ResponseWriter, r *http.Request) {
	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	var req cartitem.BulkUpdateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.BulkUpdateCart(buyerID, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.SendData(w, map[string]string{"message": "Cart updated"}, 200)
}
