package sellerdashboardhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// GetOrderDetails - GET /api/seller/dashboard/orders/{id}
func (h *DashboardHandler) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	userRoleVal := r.Context().Value(constants.ContextKeyRole)
	if userIDVal == nil || userRoleVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	regUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get order ID from path
	orderIDStr := r.PathValue("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := h.DashboardService.GetOrderDetails(regUserID, uint(orderID))
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, order, "Order details retrieved successfully")
}
