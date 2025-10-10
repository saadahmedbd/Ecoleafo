package carthandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CartHandler) GetCartTotal(w http.ResponseWriter, r *http.Request) {
	buyerID, err := getBuyerID(w, r)
	if err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	total, err := h.service.GetCartTotal(buyerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]float64{"total": total}, 200)
}
