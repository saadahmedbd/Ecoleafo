package orderhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *OrderHandler) GetOrderByNumber(w http.ResponseWriter, r *http.Request) {
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
	orderNumber := r.URL.Query().Get("order_number")
	if orderNumber == "" {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	order, err := h.orderService.GetOrderByOrderNumber(orderNumber, uint(userID), userIDType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, order, 200)

}
