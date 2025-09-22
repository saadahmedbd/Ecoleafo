package profilehandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

type BuyerProfile struct {
	UserId uint `json:"user_id"`
	// Add other profile fields as needed
	Phone          string `json:"phone"`
	DefaultAddress string `json:"default_address"`
}

func CompleteBuyerProfile(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to complete buyer profile
	var buyerProfile BuyerProfile
	encode := json.NewDecoder(r.Body)
	err := encode.Decode(&buyerProfile)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	var existingBuyerProfile models.Buyer
	if err := Config.DB.Where("user_id= ?", buyerProfile.UserId).First(&existingBuyerProfile).Error; err == nil {
		http.Error(w, "user profile already exists", http.StatusNotFound)
		return
	}
	newBuyerProfile := models.Buyer{
		RoleID:         3,
		UserId:         buyerProfile.UserId,
		Phone:          buyerProfile.Phone,
		DefaultAddress: buyerProfile.DefaultAddress,
	}
	if err := Config.DB.Create(&newBuyerProfile).Error; err != nil {
		http.Error(w, "Failed to create new Buyer profile", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]interface{}{
		"message": " user profile completed sucessfully",
		"user_id": newBuyerProfile.ID,
	}, http.StatusAccepted)

}
