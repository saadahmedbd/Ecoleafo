package selleraccounthandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerRegistrationHandler) SetDefaultPaymentMethod(w http.ResponseWriter, r *http.Request) {
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
	paymentMethodID, err := strconv.Atoi(paymentMethodIDstr)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	if err := h.service.SetDefaultPaymentMethod(uint(userID), uint(paymentMethodID)); err != nil {
		http.Error(w, `{"error":"failed to set default payment method"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{"message": "default payment method set successfully"}
	util.SendData(w, response, http.StatusOK)

}
