package inventoryhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// GetInventoryStats - GET /api/seller/inventory/stats
func (h *Inventoryhandler) GetInventoryStats(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	userRoleVal := r.Context().Value(constants.ContextKeyRole)
	if userIDVal == nil || userRoleVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	regUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get actual seller User.ID from RegUser.ID
	sellerID, err := h.inventoryservice.GetSellerIDByRegUserID(regUserID)
	if err != nil {
		util.RespondError(w, http.StatusUnauthorized, "Seller not found")
		return
	}

	stats, err := h.inventoryservice.GetInventoryStats(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Statistics retrieved successfully")

}
