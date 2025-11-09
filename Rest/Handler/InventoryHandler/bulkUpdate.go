package inventoryhandler

import (
	"encoding/json"
	"net/http"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// BulkUpdateStock - POST /api/seller/inventory/bulk-update
func (h *Inventoryhandler) BulkUpdateStock(w http.ResponseWriter, r *http.Request) {
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

	// Parse request body
	var req inventory.BulkUpdateStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Bulk update
	result, err := h.inventoryservice.BulkUpdateStock(sellerID, req.Updates, uint(sellerID))
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, result, "Bulk update successfully")

}
