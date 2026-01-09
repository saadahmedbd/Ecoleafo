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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Check if refresh token is provided and not 'undefined'
	if req.RefreshToken == "" || req.RefreshToken == "undefined" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Refresh token is required"})
		return
	}

	// Validate refresh token
	refreshToken, err := util.ValidateRefreshToken(Config.DB, req.RefreshToken)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid or expired refresh token"})
		return
	}

	// Get user details
	var user models.RegUser
	if err := Config.DB.First(&user, refreshToken.UserID).Error; err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	// Revoke old refresh token
	if err := util.RevokeRefreshToken(Config.DB, req.RefreshToken); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to revoke old token"})
		return
	}

	// Get user roles
	roles := []string{user.Role}

	// Create new token pair
	tokenPair, err := util.CreateTokenPair(Config.DB, user.ID, user.FirstName, user.LastName, roles, user.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create tokens"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenPair)
}
