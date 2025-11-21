package commissionpayoutearningshandler

import (
	"encoding/json"
	"net/http"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// POST /api/admin/payouts/request (Can be called by seller)
func (h *CommissionHandler) RequestPayout(w http.ResponseWriter, r *http.Request) {
	var req commissionearningpayoutdto.CreatePayoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Get seller ID from context
	ctxUserID := r.Context().Value(constants.ContextKeyUserID)
	ctxRoles := r.Context().Value(constants.ContextKeyRole)

	if ctxUserID == nil || ctxRoles == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Convert roles
	roleList, ok := ctxRoles.([]interface{})
	if !ok {
		http.Error(w, "Invalid role type", http.StatusInternalServerError)
		return
	}

	// Check if buyer
	isAdmin := false
	for _, raw := range roleList {
		if r, ok := raw.(string); ok && r == "seller" {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		http.Error(w, "Only seller can access this", http.StatusForbidden)
		return
	}

	// Convert user id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	// Convert reguser.id → buyer.id
	sellerID, err := getSellerID(uid)
	req.SellerID = sellerID
	if err != nil {
		http.Error(w, `{"error":"seller account not found"}`, http.StatusNotFound)
		return
	}

	if err := h.commissionService.RequestPayout(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusCreated, nil, "Payout request submitted successfully")
}
