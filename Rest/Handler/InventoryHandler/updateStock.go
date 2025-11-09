package inventoryhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// UpdateStock - PUT /api/seller/inventory/{productId}/stock
func (h *Inventoryhandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
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

	// Parse request body
	var req inventory.UpdateStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Update stock
	product, err := h.inventoryservice.UpdateStock(sellerID, uint(productID), req.Quantity, sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, product, "stock update successfully")

}
