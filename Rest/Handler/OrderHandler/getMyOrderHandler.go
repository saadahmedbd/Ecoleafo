package orderhandler

import (
	"net/http"
	"strconv"

	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *OrderHandler) GetMyOrders(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userIDType := r.Header.Get("user_role")

	if userIDStr == "" || userIDType == "" {
		http.Error(w, "user_id or user_role is empty", http.StatusBadRequest)
		return
	}

	//convert user id to uint
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return

	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	var orders []order.OrderResponse
	var total int64

	if userIDType == "[buyer]" {
		buyerID, err := getBuyerID(w, r)
		if err != nil {
			http.Error(w, "Buyer not found", http.StatusNotFound)
			return
		}
		orders, total, err = h.orderService.GetBuyerOrders(buyerID, page, limit)
	} else if userIDType == "[seller]" {
		orders, total, err = h.orderService.GetSellerOrders(uint(userID), page, limit)
	} else {
		http.Error(w, "invalid user role", http.StatusBadRequest)
		return
	}
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
