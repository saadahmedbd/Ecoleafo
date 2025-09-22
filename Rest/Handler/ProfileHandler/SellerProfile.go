package profilehandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

type SellerProfile struct {
	UserId uint `json:"user_id"`
	// Add other profile fields as needed
	BusinessEmail string `json:"business_email"`
}

func BecomeSeller(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to complete seller profile
	var sellerProfile SellerProfile
	encode := json.NewDecoder(r.Body)
	err := encode.Decode(&sellerProfile)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	// Check if seller profile already exists
	// If not, create a new seller profile
	var existingSeller models.User
	if err := Config.DB.Where("user_id= ?", sellerProfile.UserId).First(&existingSeller).Error; err == nil {
		http.Error(w, "Seller profile already exists", http.StatusNotFound)
		return
	}
	newSellerProfile := models.User{
		UserId:        sellerProfile.UserId,
		BusinessEmail: sellerProfile.BusinessEmail,
	}
	if err != Config.DB.Create(&newSellerProfile).Error {
		http.Error(w, "Error creating seller profile", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]interface{}{
		"message": " seller profile completed sucessfully",
		"user_id": newSellerProfile.ID,
	}, http.StatusCreated)

}
