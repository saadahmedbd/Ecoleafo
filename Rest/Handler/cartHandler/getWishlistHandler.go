package carthandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) GetWishlist(w http.ResponseWriter, r *http.Request) {
	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	wishlist, err := h.service.GetWishlist(buyerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, wishlist, 200)
}
