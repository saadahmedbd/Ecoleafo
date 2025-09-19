package authhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
	"gorm.io/gorm"
)

type RegistationReq struct {
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`

	// role is buyer and seller
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	///for buyer
	Phone   string `json:"phone" gorm:"not null"`
	Address string `json:"address,omitempty"` // For buyers
	//for seller
	StoreName     string `json:"store_name,omitempty"`     // For sellers
	BusinessEmail string `json:"business_email,omitempty"` // For sellers`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func Registation(w http.ResponseWriter, r *http.Request) {
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
	// req.Role = strings.ToLower(strings.TrimSpace(req.Role))
	// check user valid input
	if req.Email == "" || req.Password == "" || req.Role == "" {
		http.Error(w, "please provide email, password and user role(buyer or seller)", http.StatusBadRequest)
		return
	}
	// check role
	if !util.ValidRole(req.Role) {
		http.Error(w, "please provide valid role seller or buyer", http.StatusBadRequest)
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

	// Validate role-specific fields before any database operations
	switch req.Role {
	case "buyer":
		if req.Address == "" || req.Phone == "" {
			http.Error(w, "missing buyer field", http.StatusBadRequest)
			return
		}
	case "seller":
		if req.StoreName == "" || req.BusinessEmail == "" {
			http.Error(w, "missing seller fields", http.StatusBadRequest)
			return
		}
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
		Email:    req.Email,
		Password: hashed,
		Role:     req.Role,
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	switch req.Role {
	case "buyer":
		buyer := models.Buyer{
			RoleID:         3,
			UserId:         user.ID,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Email:          req.Email,
			Password:       hashed,
			DefaultAddress: req.Address,
			Phone:          req.Phone,
		}
		if err := tx.Create(&buyer).Error; err != nil {
			tx.Rollback()
			http.Error(w, "error creating buyer: "+err.Error(), http.StatusInternalServerError)
			return
		}
	case "seller":
		seller := models.User{
			RoleID:    2,
			UserId:    user.ID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Phone:     req.Phone,
			StoreName: req.StoreName,
			Email:     req.Email,
			Password:  hashed,
		}
		if err := tx.Create(&seller).Error; err != nil {
			tx.Rollback()
			http.Error(w, "error creating seller: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	tx.Commit()
	util.SendData(w, map[string]interface{}{
		"message": req.Role + " registered successfully",
	}, http.StatusCreated)

	}