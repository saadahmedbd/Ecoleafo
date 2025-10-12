package orderhandler

import (
	"net/http"
	"strconv"

	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	userIDType := r.Header.Get("user_role")
	if userIDType != "[admin]" {
		http.Error(w, "Admin access required", http.StatusBadRequest)
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

	orders, total, err := h.orderService.GetAllOrders(filter)
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
