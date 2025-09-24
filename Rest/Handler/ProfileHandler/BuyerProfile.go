package profilehandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

type BuyerProfile struct {

	// Add other profile fields as needed
	Phone          string `json:"phone"`
	DefaultAddress string `json:"default_address"`
}

func (h *Handler) CompleteBuyerProfile(w http.ResponseWriter, r *http.Request) {
	claimsRaw := r.Context().Value("claims")
	if claimsRaw == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	claims := claimsRaw.(map[string]interface{})
	uidFloat, ok := claims["user_id"].(float64)
	if !ok {
		// try int
		if uidInt, ok2 := claims["user_id"].(int64); ok2 {
			uidFloat = float64(uidInt)
		} else {
			http.Error(w, "invalid user id in token", http.StatusUnauthorized)
			return
		}
	}
	userID := uint(uidFloat)

	// Implement the logic to complete buyer profile
	var buyerProfile BuyerProfile
	encode := json.NewDecoder(r.Body)
	err := encode.Decode(&buyerProfile)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	var existingBuyerProfile models.Buyer
	if err := Config.DB.Where("user_id= ?", userID).First(&existingBuyerProfile).Error; err == nil {
		http.Error(w, "user profile already exists", http.StatusNotFound)
		return
	}
	newBuyerProfile := models.Buyer{
		RoleID:         3,
		UserId:         userID,
		Phone:          buyerProfile.Phone,
		DefaultAddress: buyerProfile.DefaultAddress,
	}
	if err := Config.DB.Create(&newBuyerProfile).Error; err != nil {
		http.Error(w, "Failed to create new Buyer profile", http.StatusInternalServerError)
		return
	}
	//update reguser role field
	var user models.RegUser
	if err := Config.DB.First(&user, userID).Error; err == nil {
		if !strings.Contains(user.Role, "buyer") {
			if user.Role != "" {
				user.Role += ",buyer"
			} else {
				user.Role = "buyer"
			}
		}
		Config.DB.Save(&user)
	}
	util.SendData(w, map[string]interface{}{
		"message": " user profile completed sucessfully",
		"user_id": newBuyerProfile.ID,
	}, http.StatusAccepted)

}
