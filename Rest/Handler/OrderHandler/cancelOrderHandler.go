package orderhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// for buyer
func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	// Get user data from context
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
	isBuyer := false
	for _, raw := range roleList {
		if r, ok := raw.(string); ok && r == "buyer" {
			isBuyer = true
			break
		}
	}

	if !isBuyer {
		http.Error(w, "Only buyers can access this", http.StatusForbidden)
		return
	}

	// Convert user id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	// Convert reguser.id → buyer.id
	buyerID, err := getBuyerID(uid)
	if err != nil {
		http.Error(w, `{"error":"Buyer account not found"}`, http.StatusNotFound)
		return
	}

	//take first name last name by jwt token
	claims, _ := r.Context().Value("claims").(map[string]interface{})
	firstName, _ := claims["first_name"].(string)
	lastName, _ := claims["last_name"].(string)
	username := firstName + " " + lastName

	orderId := r.URL.Query().Get("id")
	if orderId == "" {
		http.Error(w, "order id is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(orderId)
	if err != nil {
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}
	order, err := h.orderService.CancelOrder(uint(id), uint(buyerID), "buyer", username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.SendData(w, order, 200)

}
