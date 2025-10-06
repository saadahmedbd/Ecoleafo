package guestcarthandler

import (
	"fmt"
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/cart
func (h *GuestcartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	sessionID := h.getOrCreateSessionID(r, w)
	userID := h.getUserID(r)

	cart, err := h.service.GetCart(sessionID, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	util.SendData(w, cart, http.StatusOK)
}
