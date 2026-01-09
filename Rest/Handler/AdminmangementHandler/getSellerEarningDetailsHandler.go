package adminmangementhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *AdminManagementHandler) GetSellerEarningDetails(w http.ResponseWriter, r *http.Request) {
	sellerIDStr := r.PathValue("id")
	if sellerIDStr == "" {
		util.SendError(w, "seller_id is required", http.StatusBadRequest)
		return
	}

	sellerID, err := strconv.ParseUint(sellerIDStr, 10, 32)
	if err != nil {
		util.SendError(w, "invalid seller_id", http.StatusBadRequest)
		return
	}

	details, err := h.commissionService.GetSellerEarningDetails(uint(sellerID))
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, details, http.StatusOK)
}
