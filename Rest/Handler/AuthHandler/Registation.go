package authhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"time"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
	"gorm.io/gorm"
)

type RegistationReq struct {

	//basic registation
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (h *Handler) Registation(w http.ResponseWriter, r *http.Request) {
	var req RegistationReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if !emailRegex.MatchString(req.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// check existing email
	var existing models.RegUser
	if err := Config.DB.Where("email=?", req.Email).First(&existing).Error; err == nil {
		http.Error(w, "email already registered", http.StatusConflict)
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	// check length of password
	if len(req.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}
	//hashed password store in database
	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "password not convert hased", http.StatusInternalServerError)
		return
	}
	// Use transaction to ensure atomicity
	tx := Config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//email, password and role save in reqUser field
	user := models.RegUser{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Password:   hashed,
		Role:       "buyer",
		Created_At: time.Now(),
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Automatically create buyer record for cart functionality
	buyer := models.Buyer{
		RoleID:   3,
		UserId:   user.ID,
		Password: hashed,
	}
	if err := tx.Create(&buyer).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error creating buyer profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	tokenPair, err := util.CreateTokenPair(Config.DB, user.ID, user.FirstName, user.LastName, []string{"buyer"}, buyer.ID)
	if err != nil {
		http.Error(w, "Failed to create tokens", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]interface{}{
		"access_token": tokenPair,
		"user_type":    user.Role,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"email":        user.Email,
		"roles":        []string{user.Role},
		"user_id":      user.ID,
	}, http.StatusCreated)

}
