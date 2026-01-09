package sellerproflehandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerProfileHandler) ChangeSellerPassword(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userIDType := r.Header.Get("user_role")

	if userIDStr == "" || userIDType == "" {
		http.Error(w, "user_id or user_role is empty", http.StatusBadRequest)
		return
	}
	if !strings.Contains(userIDType, "seller") {
		http.Error(w, "only seller can acces this account", http.StatusForbidden)
		return
	}
	//convert user id to uint
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return

	}
	var req sellerprofile.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.ChangePassword(uint(userID), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"message": "Password changed successfully",
	}
	util.SendData(w, response, http.StatusOK)
}
