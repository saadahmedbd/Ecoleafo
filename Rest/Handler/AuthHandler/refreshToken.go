package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshToken handler - returns new token pair
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if refresh token is provided and not 'undefined'
	if req.RefreshToken == "" || req.RefreshToken == "undefined" {
		http.Error(w, "Refresh token is required and must be valid", http.StatusBadRequest)
		return
	}

	// Validate refresh token
	refreshToken, err := util.ValidateRefreshToken(Config.DB, req.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	// Get user details
	var user models.RegUser
	if err := Config.DB.First(&user, refreshToken.UserID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Revoke old refresh token
	if err := util.RevokeRefreshToken(Config.DB, req.RefreshToken); err != nil {
		http.Error(w, "Failed to revoke old token", http.StatusInternalServerError)
		return
	}

	// Get user roles
	roles := []string{user.Role}

	// Create new token pair
	tokenPair, err := util.CreateTokenPair(Config.DB, user.ID, user.FirstName, user.LastName, roles)
	if err != nil {
		http.Error(w, "Failed to create tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenPair)
}
