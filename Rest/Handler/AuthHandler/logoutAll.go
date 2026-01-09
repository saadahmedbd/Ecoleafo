package authhandler

import (
	"net/http"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	util "github.com/saadahmedbd/Treestore/Util"
)

// LogoutAll handler - revokes all user tokens
func (h *Handler) LogoutAll(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from token
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := util.VerifyJwt(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID := uint(claims["user_id"].(float64))

	if err := util.RevokeAllUserTokens(Config.DB, userID); err != nil {
		http.Error(w, "Failed to logout from all devices", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Logged out from all devices", 200)
}
