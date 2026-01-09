package guestcarthandler

import (
	"fmt"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// POST /api/cart/migrate (Protected - Called after login)
func (h *GuestcartHandler) MigrateGuestCart(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	if userIDStr == "" {
		http.Error(w, `{"error":"authentication required"}`, http.StatusUnauthorized)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}

	// Get session ID from cookie
	cookie, err := r.Cookie("guest_session")
	if err != nil || cookie.Value == "" {
		http.Error(w, `{"error":"no guest session found"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.MigrateGuestCart(cookie.Value, uint(userID)); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Clear guest session cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "guest_session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	response := map[string]interface{}{
		"message": "Cart migrated successfully",
	}
	util.SendData(w, response, http.StatusOK)
}
