package authhandler

import (
	"encoding/json"
	"errors"

	"net/http"
	"time"

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
	Token     string   `json:"token"`
	UserType  string   `json:"user_type"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
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

	// Try to find in buyers table
	var buyer models.Buyer
	buyerErr := h.service.db.Where("user_id = ?", regUser.ID).First(&buyer).Error
	if buyerErr == nil {
		// Found in buyers table
		userType = "buyer"
		userRoles = []string{"buyer"}
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

	}

	// If not found in either table, check if this is an admin
	if buyerErr != nil && sellerErr != nil {
		// Assume admin or default buyer
		userRoles = []string{"buyer"}
		userType = "buyer"
	}

	// Generate JWT token with the RegUser.ID (which is referenced by user_id in other tables)
	token, err := util.CreateJwt(regUser.ID, firstName, lastName, userRoles, 24*time.Hour)
	if err != nil {
		http.Error(w, `{"error":"failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := LoginResponse{
		Token:     token,
		UserType:  userType,
		FirstName: firstName,
		LastName:  lastName,
		Email:     regUser.Email,
		Roles:     userRoles,
	}

	util.SendData(w, response, http.StatusOK)
}

// type Reqlogin struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
// 	var req Reqlogin
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
// 		return
// 	}

// 	if !emailRegex.MatchString(req.Email) {
// 		http.Error(w, `{"error":"Invalid email format"}`, http.StatusBadRequest)
// 		return
// 	}

// 	var user models.RegUser
// 	if err := Config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
// 		http.Error(w, `{"error":"Invalid credentials"}`, http.StatusUnauthorized)
// 		return
// 	}

// 	if err := user.CheckPassword(req.Password); err != nil {
// 		http.Error(w, `{"error":"Invalid credentials"}`, http.StatusUnauthorized)
// 		return
// 	}
// 	//check roles
// 	roles := []string{}
// 	//check user is buyer
// 	var buyer models.Buyer
// 	if err := Config.DB.Where("user_id = ?", user.ID).First(&buyer).Error; err == nil {
// 		roles = append(roles, "buyer")
// 	}
// 	//check user is seller
// 	var seller models.User
// 	if err := Config.DB.Where("user_id = ?", user.ID).First(&seller).Error; err == nil {
// 		roles = append(roles, "seller")
// 	}

// 	// If no specific roles found, use RegUser.Role as fallback
// 	if len(roles) == 0 && user.Role != "" {
// 		roles = append(roles, user.Role)
// 	}
// 	token, err := util.CreateJwt(user.ID, user.FirstName, user.LastName, roles, 24*time.Hour)
// 	if err != nil {
// 		http.Error(w, "Error generating token", http.StatusInternalServerError)
// 		return
// 	}

// 	// cnf := Config.GetConfig()
// 	// token, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
// 	// 	UserId:    user.ID,
// 	// 	Role:      user.Role,
// 	// 	FirstName: user.FirstName,
// 	// 	LastName:  user.LastName,
// 	// })
// 	// if err != nil {
// 	// 	http.Error(w, "Error generating token", http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	w.Header().Set("Content-Type", "application/json")
// 	response := map[string]interface{}{
// 		"message":    "Login successful",
// 		"email":      user.Email,
// 		"user id":    user.ID,
// 		"token":      token,
// 		"role":       user.Role,
// 		"first_name": user.FirstName,
// 		"last_name":  user.LastName,
// 	}

// 	json.NewEncoder(w).Encode(response)
// }
