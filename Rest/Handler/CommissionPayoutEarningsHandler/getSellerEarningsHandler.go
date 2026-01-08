package commissionpayoutearningshandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/sellers/earnings?seller_id=5
func (h *CommissionHandler) GetSellerEarnings(w http.ResponseWriter, r *http.Request) {
	sellerIDStr := r.URL.Query().Get("seller_id")
	userIDStr := r.URL.Query().Get("user_id")
	
	var sellerID uint
	var err error
	
	if sellerIDStr != "" {
		id, parseErr := strconv.ParseUint(sellerIDStr, 10, 64)
		if parseErr != nil {
			util.RespondError(w, http.StatusBadRequest, "Invalid seller_id")
			return
		}
		sellerID = uint(id)
	} else if userIDStr != "" {
		// Convert user_id to seller_id
		userID, parseErr := strconv.ParseUint(userIDStr, 10, 64)
		if parseErr != nil {
			util.RespondError(w, http.StatusBadRequest, "Invalid user_id")
			return
		}
		sellerID, err = h.commissionService.GetSellerIDByUserID(uint(userID))
		if err != nil {
			util.RespondError(w, http.StatusNotFound, "Seller not found")
			return
		}
	} else {
		util.RespondError(w, http.StatusBadRequest, "seller_id or user_id is required")
		return
	}

	earnings, err := h.commissionService.GetSellerEarnings(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get earnings")
		return
	}

	util.RespondJSON(w, http.StatusOK, earnings, "Seller earnings retrieved successfully")
}
