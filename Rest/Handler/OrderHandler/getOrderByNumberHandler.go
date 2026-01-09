package orderhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

func (h *OrderHandler) GetOrderByNumber(w http.ResponseWriter, r *http.Request) {
	//Get user data from context
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
		if roleStr, ok := raw.(string); ok && roleStr == "buyer" {
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
	orderNumber := r.URL.Query().Get("order_number")
	if orderNumber == "" {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	order, err := h.orderService.GetOrderByOrderNumber(orderNumber, uint(buyerID), "buyer")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, order, 200)

}
