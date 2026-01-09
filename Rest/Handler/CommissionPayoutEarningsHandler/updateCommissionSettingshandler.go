package commissionpayoutearningshandler

import (
	"encoding/json"
	"net/http"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// PUT /api/admin/commission/settings
func (h *CommissionHandler) UpdateCommissionSettings(w http.ResponseWriter, r *http.Request) {
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

	var req commissionearningpayoutdto.UpdateCommissionSettingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.commissionService.UpdateCommissionSettings(&req, adminID); err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to update settings")
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Commission settings updated successfully")
}
