package commissionpayoutearningshandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// POST /api/admin/payouts/process?id=5
func (h *CommissionHandler) ProcessPayout(w http.ResponseWriter, r *http.Request) {
	payoutIDStr := r.URL.Query().Get("id")
	if payoutIDStr == "" {
		util.RespondError(w, http.StatusBadRequest, "payout id is required")
		return
	}

	payoutID, err := strconv.ParseUint(payoutIDStr, 10, 64)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid payout id")
		return
	}

	var req commissionearningpayoutdto.ProcessPayoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Get admin ID from context
	// Get admin ID from context (set by auth middleware)
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
		if r, ok := raw.(string); ok && r == "admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		http.Error(w, "Only admin can access this", http.StatusForbidden)
		return
	}

	// Convert user id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	// Convert reguser.id → buyer.id
	adminID, err := getAdminID(uid)
	if err != nil {
		http.Error(w, `{"error":"Buyer account not found"}`, http.StatusNotFound)
		return
	}

	if err := h.commissionService.ProcessPayout(uint(payoutID), &req, adminID); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Payout processed successfully")
}
