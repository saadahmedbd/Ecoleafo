package sellerproflehandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerProfileHandler) UpdateStoreImage(w http.ResponseWriter, r *http.Request) {
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
	var req sellerprofile.UpdateStoreImageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateStoreImages(uint(userID), req); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Store images updated successfully",
	}
	util.SendData(w, response, http.StatusOK)
}
