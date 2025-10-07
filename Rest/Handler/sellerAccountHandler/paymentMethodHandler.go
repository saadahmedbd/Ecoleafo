package selleraccounthandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerRegistrationHandler) AddPaymentMethod(w http.ResponseWriter, r *http.Request) {
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
	var req selleraccount.AddPaymentMethodRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	paymentMethod, err := h.service.AddPaymentMethod(uint(userID), &req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	util.SendData(w, paymentMethod, http.StatusOK)

}
