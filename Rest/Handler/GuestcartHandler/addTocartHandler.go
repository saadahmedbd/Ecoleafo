package guestcarthandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	guestcartitem "github.com/saadahmedbd/Treestore/Rest/DTO/GuestCartItem"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *GuestcartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	sessionID := h.getOrCreateSessionID(r, w)
	userID := h.getUserID(r)

	var req guestcartitem.AddToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	cart, err := h.service.AddToCart(sessionID, userID, req)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	util.SendData(w, cart, http.StatusOK)

}
