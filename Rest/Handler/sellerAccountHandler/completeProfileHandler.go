package selleraccounthandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerRegistrationHandler) CompleteProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	if userIDStr == "" || userType == "" {
		util.SendError(w, "authentication required", http.StatusUnauthorized)
		return
	}

	if !strings.Contains(userType, "seller") {
		util.SendError(w, "only sellers can access this", http.StatusForbidden)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		util.SendError(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	
	var req selleraccount.CompleteSellerProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	
	profile, err := h.service.CompleteSellerProfile(uint(userID), &req)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	util.SendData(w, profile, http.StatusOK)
}
