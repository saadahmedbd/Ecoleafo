package profilehandler

import (
	"encoding/json"
	"net/http"

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

	var buyerProfile BuyerProfile
	encode := json.NewDecoder(r.Body)
	err := encode.Decode(&buyerProfile)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	var existingBuyerProfile models.Buyer
	if err := Config.DB.Where("user_id= ?", userID).First(&existingBuyerProfile).Error; err != nil {
		http.Error(w, "buyer profile not found", http.StatusNotFound)
		return
	}
	existingBuyerProfile.Phone = buyerProfile.Phone
	existingBuyerProfile.DefaultAddress = buyerProfile.DefaultAddress
	if err := Config.DB.Save(&existingBuyerProfile).Error; err != nil {
		http.Error(w, "Failed to update buyer profile", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]interface{}{
		"message": "buyer profile updated successfully",
		"user_id": existingBuyerProfile.ID,
	}, http.StatusAccepted)

}
