package orderhandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userIDType := r.Header.Get("user_role")

	if userIDStr == "" || userIDType == "" {
		http.Error(w, "user_id or user_role is empty", http.StatusBadRequest)
		return
	}
	if !strings.Contains(userIDType, "buyer") {
		http.Error(w, "only buyer can create order", http.StatusForbidden)
		return
	}
	//convert user id to uint
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	// Get buyer ID from user ID
	var buyer models.Buyer
	if err := Config.DB.Where("user_id = ?", uint(userID)).First(&buyer).Error; err != nil {
		http.Error(w, "Buyer not found", http.StatusNotFound)
		return
	}

	var req order.CreateOrderRequest
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&req)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	order, err := h.orderService.CreateOrder(buyer.ID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.SendData(w, order, 200)

}
