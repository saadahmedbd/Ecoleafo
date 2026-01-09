package inventoryhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// GetStockHistory - GET /api/seller/inventory/{productId}/history
func (h *Inventoryhandler) GetStockHistory(w http.ResponseWriter, r *http.Request) {
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

	// Get product ID from path
	productIDStr := r.PathValue("productId")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Get limit
	limit := 10
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	// Get history
	history, err := h.inventoryservice.GetStockHistory(sellerID, uint(productID), limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]interface{}{"history": history, "count": len(history)}, "Inventory retrieved successfully")
}
