package orderhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *OrderHandler) GetOrderHistory(w http.ResponseWriter, r *http.Request) {
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

	orderId := r.URL.Query().Get("id")
	if orderId == "" {
		http.Error(w, "order id is empty", http.StatusBadRequest)
		return
	}
	sId, err := strconv.Atoi(orderId)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusNoContent)
		return
	}
	history, err := h.orderService.GetOrderHistory(uint(sId), uint(userID), userIDType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, history, 200)
}
