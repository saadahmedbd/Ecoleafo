package authhandler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
	"gorm.io/gorm"
)

// GoogleAuthRequest contains the role for registration
type GoogleAuthRequest struct {
	Role string `json:"role"` // "buyer" or "seller"
}

// GoogleAuthResponse is the response after successful OAuth
type GoogleAuthResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int64    `json:"expires_in"`
	UserType     string   `json:"user_type"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
	IsNewUser    bool     `json:"is_new_user"`
	NeedsProfile bool     `json:"needs_profile"` // For sellers who need to complete profile
}

// GoogleLoginInit initiates the Google OAuth flow
// Endpoint: GET /auth/google/login?role=buyer or GET /auth/google/login?role=seller
func (h *Handler) GoogleLoginInit(w http.ResponseWriter, r *http.Request) {
	// Get role from query parameter (buyer or seller)
	role := r.URL.Query().Get("role")
	if role == "" {
		role = "buyer" // Default to buyer
	}

	// Validate role
	if role != "buyer" && role != "seller" {
		http.Error(w, `{"error":"invalid role. Must be 'buyer' or 'seller'"}`, http.StatusBadRequest)
		return
	}

	// Generate state token for CSRF protection
	stateToken, err := generateStateToken()
	if err != nil {
		http.Error(w, `{"error":"failed to generate state"}`, http.StatusInternalServerError)
		return
	}

	// Store state and role in database with expiry (industry standard for cross-domain)
	oauthState := models.OAuthState{
		State:     stateToken,
		Role:      role,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	h.service.db.Create(&oauthState)

	// Get OAuth config and generate auth URL
	oauthConfig := util.GetGoogleOAuthConfig()
	url := oauthConfig.AuthCodeURL(stateToken)

	// Redirect to Google's OAuth consent page
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GoogleCallback handles the callback from Google OAuth
// Endpoint: GET /auth/google/callback
func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Get state from query parameter
	stateParam := r.URL.Query().Get("state")
	if stateParam == "" {
		h.redirectToFrontendWithError(w, r, "missing state parameter")
		return
	}

	// Retrieve state from database
	var oauthState models.OAuthState
	err := h.service.db.Where("state = ? AND expires_at > ?", stateParam, time.Now()).First(&oauthState).Error
	if err != nil {
		h.redirectToFrontendWithError(w, r, "invalid or expired state")
		return
	}

	role := oauthState.Role

	// Delete used state
	h.service.db.Delete(&oauthState)

	// Get authorization code
	code := r.URL.Query().Get("code")
	if code == "" {
		h.redirectToFrontendWithError(w, r, "missing authorization code")
		return
	}

	// Exchange code for token
	oauthConfig := util.GetGoogleOAuthConfig()
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		h.redirectToFrontendWithError(w, r, "failed to exchange token")
		return
	}

	// Get user info from Google
	googleUser, err := util.GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		h.redirectToFrontendWithError(w, r, "failed to get user info")
		return
	}

	// Process the Google user and create/login
	response, err := h.processGoogleUser(googleUser, role)
	if err != nil {
		h.redirectToFrontendWithError(w, r, err.Error())
		return
	}

	// Redirect to frontend with tokens
	h.redirectToFrontendWithSuccess(w, r, response)
}

// processGoogleUser handles the logic for creating or logging in a Google user
func (h *Handler) processGoogleUser(googleUser *util.GoogleUserInfo, role string) (*GoogleAuthResponse, error) {
	var regUser models.RegUser
	var isNewUser bool

	// Check if user exists by GoogleID
	err := h.service.db.Where("google_id = ?", googleUser.ID).First(&regUser).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User doesn't exist by GoogleID, check by email
			err = h.service.db.Where("email = ?", googleUser.Email).First(&regUser).Error

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// Completely new user - create account
					regUser, err = h.createGoogleUser(googleUser, role)
					if err != nil {
						return nil, err
					}
					isNewUser = true
				} else {
					return nil, fmt.Errorf("database error: %w", err)
				}
			} else {
				// User exists with this email but different auth method
				// Link Google account to existing account
				if regUser.GoogleID == "" {
					regUser.GoogleID = googleUser.ID
					regUser.AuthProvider = "google"
					regUser.EmailVerified = true
					regUser.IsVerified = true
					if regUser.Avatar == "" && googleUser.Picture != "" {
						regUser.Avatar = googleUser.Picture
					}
					h.service.db.Save(&regUser)
				} else {
					// Email exists but with different GoogleID - shouldn't happen
					return nil, fmt.Errorf("email already registered with different Google account")
				}
			}
		} else {
			return nil, fmt.Errorf("database error: %w", err)
		}
	}

	// Update last login
	now := time.Now()
	regUser.LastLoginAt = &now
	h.service.db.Model(&regUser).Update("last_login_at", now)

	// Determine user type and create profile if needed
	userType, needsProfile, err := h.ensureUserProfile(&regUser, role, isNewUser)
	if err != nil {
		return nil, err
	}

	// Generate tokens
	tokenPair, err := util.CreateTokenPair(h.service.db, regUser.ID, regUser.FirstName, regUser.LastName, []string{regUser.Role}, regUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create tokens: %w", err)
	}

	// Build response
	response := &GoogleAuthResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresIn:    tokenPair.ExpiresIn,
		UserType:     userType,
		FirstName:    regUser.FirstName,
		LastName:     regUser.LastName,
		Email:        regUser.Email,
		Roles:        []string{regUser.Role},
		IsNewUser:    isNewUser,
		NeedsProfile: needsProfile,
	}

	return response, nil
}

// createGoogleUser creates a new user from Google OAuth data
func (h *Handler) createGoogleUser(googleUser *util.GoogleUserInfo, role string) (models.RegUser, error) {
	// Validate role
	if role != "buyer" && role != "seller" {
		role = "buyer" // Default to buyer
	}

	// Create RegUser
	regUser := models.RegUser{
		FirstName:     googleUser.GivenName,
		LastName:      googleUser.FamilyName,
		Email:         googleUser.Email,
		GoogleID:      googleUser.ID,
		AuthProvider:  "google",
		Role:          role,
		Phone:         "N/A",
		Avatar:        googleUser.Picture,
		IsActive:      true,
		IsVerified:    true, // Google verified
		EmailVerified: googleUser.VerifiedEmail,
		Password:      "", // No password for OAuth users
	}

	if err := h.service.db.Create(&regUser).Error; err != nil {
		return models.RegUser{}, fmt.Errorf("failed to create user: %w", err)
	}

	return regUser, nil
}

// ensureUserProfile creates buyer/seller profile if needed
func (h *Handler) ensureUserProfile(regUser *models.RegUser, requestedRole string, isNewUser bool) (string, bool, error) {
	var userType string
	var needsProfile bool

	// For existing users, use their current role
	// For new users, use the requested role
	effectiveRole := regUser.Role
	if isNewUser {
		effectiveRole = requestedRole
	}

	switch effectiveRole {
	case "buyer":
		// Check if buyer profile exists
		var buyer models.Buyer
		err := h.service.db.Where("user_id = ?", regUser.ID).First(&buyer).Error

		if errors.Is(err, gorm.ErrRecordNotFound) && isNewUser {
			// Create buyer profile with default role
			var buyerRole models.Role
			h.service.db.Where("name = ?", "buyer").First(&buyerRole)

			buyer = models.Buyer{
				UserId:        regUser.ID,
				RoleID:        buyerRole.ID,
				Password:      "", // OAuth users don't need password in buyer table
				Phone:         regUser.Phone,
				IsActive:      true,
				EmailVerified: regUser.EmailVerified,
			}

			if err := h.service.db.Create(&buyer).Error; err != nil {
				return "", false, fmt.Errorf("failed to create buyer profile: %w", err)
			}
		}

		userType = "buyer"
		needsProfile = false // Buyers don't need additional profile setup

	case "seller":
		// Check if seller profile exists
		var seller models.User
		err := h.service.db.Where("user_id = ?", regUser.ID).First(&seller).Error

		if errors.Is(err, gorm.ErrRecordNotFound) && isNewUser {
			// Create minimal seller profile
			var sellerRole models.Role
			h.service.db.Where("name = ?", "seller").First(&sellerRole)

			seller = models.User{
				UserId:            regUser.ID,
				RoleID:            sellerRole.ID,
				BusinessEmail:     regUser.Email,
				Password:          "", // OAuth users don't need password
				Phone:             regUser.Phone,
				StoreName:         fmt.Sprintf("%s's Store", regUser.FirstName),
				StoreSlug:         generateStoreSlug(regUser.FirstName, regUser.ID),
				Address:           "N/A",
				Country:           "Bangladesh",
				Status:            "pending",
				ApprovalStatus:    "pending",
				IsActive:          true,
				IsProfileComplete: false,
				NextStep:          "complete_profile",
			}

			if err := h.service.db.Create(&seller).Error; err != nil {
				return "", false, fmt.Errorf("failed to create seller profile: %w", err)
			}

			needsProfile = true // Sellers need to complete their profile
		} else if err == nil {
			// Existing seller - check if profile is complete
			needsProfile = !seller.IsProfileComplete
		}

		userType = "seller"
	}

	return userType, needsProfile, nil
}

// Helper functions
func generateStateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func generateStoreSlug(name string, userID uint) string {
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	return fmt.Sprintf("%s-%d", slug, userID)
}

func (h *Handler) redirectToFrontendWithError(w http.ResponseWriter, r *http.Request, errorMsg string) {
	cnf := Config.GetConfig()
	redirectURL := fmt.Sprintf("%s/auth/callback?error=%s", cnf.FrontendURL, errorMsg)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func (h *Handler) redirectToFrontendWithSuccess(w http.ResponseWriter, r *http.Request, response *GoogleAuthResponse) {
	cnf := Config.GetConfig()

	// Convert response to JSON for URL parameter
	jsonData, _ := json.Marshal(response)
	encodedData := base64.URLEncoding.EncodeToString(jsonData)

	redirectURL := fmt.Sprintf("%s/auth/callback?data=%s", cnf.FrontendURL, encodedData)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}
