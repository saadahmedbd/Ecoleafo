package orderhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

func (h *OrderHandler) GetSellerOrders(w http.ResponseWriter, r *http.Request) {
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
		if r, ok := raw.(string); ok && r == "seller" {
			isBuyer = true
			break
		}
	}

	if !isBuyer {
		http.Error(w, "Only seller can access this", http.StatusForbidden)
		return
	}

	// Convert user id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	// Convert reguser.id → buyer.id
	sellerID, err := getSellerID(uid)
	if err != nil {
		http.Error(w, `{"error":"seller account not found"}`, http.StatusNotFound)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	orders, total, err := h.orderService.GetSellerOrders(uint(sellerID), page, limit)
	if err != nil {
		http.Error(w, "failed to get orders", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"orders": orders,
		"total":  total,
		"page":   page,
		"limit":  limit,
	}
	util.SendData(w, response, 200)
}
