package orderhandler

import (
	"net/http"
	"strconv"

	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
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
		if r, ok := raw.(string); ok && r == "admin" {
			isBuyer = true
			break
		}
	}

	if !isBuyer {
		http.Error(w, "Only admin can access this", http.StatusForbidden)
		return
	}

	// Convert user id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	// Convert reguser.id → buyer.id
	adminID, err := getAdminID(uid)
	if err != nil {
		http.Error(w, `{"error":"admin account not found"}`, http.StatusNotFound)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	status := r.URL.Query().Get("status")
	paymentStatus := r.URL.Query().Get("payment_status")
	buyerID, _ := strconv.ParseUint(r.URL.Query().Get("buyer_id"), 10, 32)

	filter := order.OrderListFilter{
		Status:        status,
		PaymentStatus: paymentStatus,
		BuyerID:       uint(buyerID),
		Page:          page,
		Limit:         limit,
	}

	orders, total, err := h.orderService.GetAllOrders(adminID, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
