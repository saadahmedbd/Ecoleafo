package sellerdashboardhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

type UpdateOrderStatusRequest struct {
	Status         string `json:"status"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	Notes          string `json:"notes,omitempty"`
}

// UpdateOrderStatus - PUT /api/seller/orders/{id}/status
func (h *DashboardHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	regUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	orderIDStr := r.PathValue("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var req UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.DashboardService.UpdateOrderStatus(regUserID, uint(orderID), req.Status, req.TrackingNumber, req.Notes); err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"message": "Order status updated successfully"}, "Order status updated")
}
