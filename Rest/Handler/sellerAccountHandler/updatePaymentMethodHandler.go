package selleraccounthandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerRegistrationHandler) UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {
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
	paymentMethodIDstr := r.PathValue("id")
	paymentMethod, err := strconv.Atoi(paymentMethodIDstr)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var req selleraccount.UpdatePaymentMethodRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	if err := h.service.UpdatePaymentMethod(uint(userID), uint(paymentMethod), &req); err != nil {
		http.Error(w, `{"error":"failed to update payment method"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{"message": "payment method updated successfully"}
	util.SendData(w, response, http.StatusOK)

}
