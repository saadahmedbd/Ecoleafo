package orderhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *OrderHandler) UpdateOrderItemStatus(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userIDType := r.Header.Get("user_role")
	//take first name last name by jwt token

	claims, _ := r.Context().Value("claims").(map[string]interface{})
	firstName, _ := claims["first_name"].(string)
	lastName, _ := claims["last_name"].(string)
	username := firstName + " " + lastName
	if username == " " {
		username = "user_" + userIDStr
	}

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
	var req order.UpdateOrderItemStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	err = h.orderService.UpdateOrderItemStatus(uint(userID), userIDType, username, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, "update order item succesfully", 200)

}
