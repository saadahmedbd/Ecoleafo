package buyerprofilehandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Buyerprofilehandler) SetDefaultAddress(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")
	if userIDStr == "" || userType == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if !strings.Contains("buyer", userType) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	addressIDStr := vars["id"]
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid address ID"}`, http.StatusBadRequest)
		return
	}

	if err := h.buyerservice.SetDefaultAddress(uint(userID), uint(addressID)); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Default address updated successfully",
	}
	util.SendData(w, response, http.StatusOK)

}
