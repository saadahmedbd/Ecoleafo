package carthandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	if err := h.service.ClearCart(buyerID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"message": "Cart cleared"}, 200)
}
