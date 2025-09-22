package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
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
	//check roles
	roles := []string{}
	//check user is buyer
	var buyer models.Buyer
	if err := Config.DB.Where("user_id = ?", user.ID).First(&buyer).Error; err == nil {
		roles = append(roles, "buyer")
	}
	//check user is seller
	var seller models.User
	if err := Config.DB.Where("user_id = ?", user.ID).First(&seller).Error; err == nil {
		roles = append(roles, "seller")
	}

	cnf := Config.GetConfig()
	token, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		UserId:    user.ID,
		Role:      user.Role,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Login successful",
		"email":   user.Email,
		"user id": user.ID,
		"token":   token,
		"role":    user.Role,
	}

	json.NewEncoder(w).Encode(response)
}
