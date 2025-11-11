package authhandler

import (
	"encoding/json"
	"errors"

	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
	"gorm.io/gorm"
)

// Supporting types
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int64    `json:"expires_in"`
	UserType     string   `json:"user_type"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		http.Error(w, `{"error":"email and password are required"}`, http.StatusBadRequest)
		return
	}

	// First, try to find the RegUser (main user record)
	var regUser models.RegUser
	if err := h.service.db.Where("email = ?", req.Email).First(&regUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, `{"error":"invalid email or password"}`, http.StatusUnauthorized)
			return
		}
		http.Error(w, `{"error":"database error"}`, http.StatusInternalServerError)
		return
	}

	// Verify password
	if err := regUser.CheckPassword(req.Password); err != nil {
		http.Error(w, `{"error":"invalid email or password"}`, http.StatusUnauthorized)
		return
	}

	// Now determine user type and get role information
	var userRoles []string
	var firstName, lastName string
	var userType string

	firstName = regUser.FirstName
	lastName = regUser.LastName

	// Update RegUser role if empty
	var needsRoleUpdate bool

	// Try to find in buyers table
	var buyer models.Buyer
	buyerErr := h.service.db.Where("user_id = ?", regUser.ID).First(&buyer).Error
	if buyerErr == nil {
		// Found in buyers table
		userType = "buyer"
		userRoles = []string{"buyer"}
		if regUser.Role == "" {
			regUser.Role = "buyer"
			needsRoleUpdate = true
		}
	}

	// Try to find in users table (sellers)
	var seller models.User
	sellerErr := h.service.db.Preload("Role").Where("user_id = ?", regUser.ID).First(&seller).Error
	if sellerErr == nil {
		// Found in sellers table
		userType = "seller"
		if seller.Role.Name != "" {
			userRoles = []string{seller.Role.Name}
		} else {
			userRoles = []string{"seller"}
		}
		if regUser.Role == "" {
			regUser.Role = "seller"
			needsRoleUpdate = true
		}
	}

	// If not found in either table, check if this is an admin
	if buyerErr != nil && sellerErr != nil {
		var admin models.Admin
		adminErr := h.service.db.Where("user_id = ?", regUser.ID).First(&admin).Error
		if adminErr == nil {
			userType = "admin"
			userRoles = []string{"admin"}
			firstName = admin.FullName
			lastName = ""
			if regUser.Role == "" {
				regUser.Role = "admin"
				needsRoleUpdate = true
			}
		} else {
			userRoles = []string{"buyer"}
			userType = "please compete buyer profile"
			if regUser.Role == "" {
				regUser.Role = "buyer"
				needsRoleUpdate = true
			}
		}
	}

	// Update RegUser role if needed
	if needsRoleUpdate {
		h.service.db.Model(&regUser).Update("role", regUser.Role)
	}

	// Create token pair have jwt and refresh token
	tokenPair, err := util.CreateTokenPair(Config.DB, regUser.ID, regUser.FirstName, regUser.LastName, userRoles)
	if err != nil {
		http.Error(w, "Failed to create tokens", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresIn:    tokenPair.ExpiresIn,
		UserType:     userType,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        regUser.Email,
		Roles:        userRoles,
	}

	util.SendData(w, response, http.StatusOK)
}
