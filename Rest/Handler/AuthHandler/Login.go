package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

type Reqlogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req Reqlogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if !emailRegex.MatchString(req.Email) {
		http.Error(w, `{"error":"Invalid email format"}`, http.StatusBadRequest)
		return
	}

	var user models.RegUser
	if err := Config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, `{"error":"Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	if err := user.CheckPassword(req.Password); err != nil {
		http.Error(w, `{"error":"Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Login successful",
		"email":   user.Email,
		"role":    user.Role,
	}

	json.NewEncoder(w).Encode(response)
}
