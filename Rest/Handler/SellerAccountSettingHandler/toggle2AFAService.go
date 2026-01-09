package selleraccountsettinghandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// TOGGLE 2FA
// PUT /api/seller/security/2fa
// ==========================================

func (h *Selleraccountsettinghandler) Toggle2FA(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	if userIDStr == "" || userType == "" {
		http.Error(w, `{"error":"authentication required"}`, http.StatusUnauthorized)
		return
	}

	if !strings.Contains(userType, "seller") {
		http.Error(w, `{"error":"only sellers can access this"}`, http.StatusForbidden)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}

	var req selleraccountsetting.Toggle2FARequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.Toggle2FA(uint(userID), req.Enabled)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "")
}
