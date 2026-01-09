package guestcarthandler

import (
	"fmt"
	"net/http"
	"time"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *GuestcartHandler) getOrCreateSessionID(r *http.Request, w http.ResponseWriter) string {
	cookie, err := r.Cookie("guest_session")
	if err == nil && cookie.Value != "" {
		return cookie.Value
	}
	//create new session id
	sessionID := fmt.Sprintf("guest_%d_%s", time.Now().Unix(), util.GenerateRandomString(16))
	http.SetCookie(w, &http.Cookie{
		Name:     "guest_session",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   86400 * 30, // 30 days
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return sessionID
}
