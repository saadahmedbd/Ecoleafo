package guestcarthandler

import (
	"net/http"
	"strconv"
)

func (h *GuestcartHandler) getUserID(r *http.Request) *uint {
	userIDStr := r.Header.Get("user_id")
	if userIDStr == "" {
		return nil
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return nil
	}

	uid := uint(userID)
	return &uid
}
