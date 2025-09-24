package profilehandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

type ComPeleteSeller struct {

	// Add other profile fields as needed
	BusinessEmail string  `json:"business_email"`
	StoreName     string  `json:"business_name"`
	Address       string  `json:"address"`
	Phone         string  `json:"phone"`
	StoreDesc     string  `json:"store_description"`
	Commission    float64 `json:"commission"`
}

func (h *Handler) CompleteSeller(w http.ResponseWriter, r *http.Request) {
	//user id bring to jwt field
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

	var completeSeller ComPeleteSeller
	encode := json.NewDecoder(r.Body)
	err := encode.Decode(&completeSeller)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	// Check if seller profile already exists
	// If not, create a new seller profile
	var seller models.RegUser
	if err := Config.DB.Where("user_id= ?", userID).First(&seller).Error; err == nil {
		http.Error(w, "user profile already exists ", http.StatusNotFound)
		return
	}
	sellerProfile := models.User{
		RoleID:        2,
		UserId:        userID,
		BusinessEmail: completeSeller.BusinessEmail,
		StoreName:     completeSeller.StoreName,
		Phone:         completeSeller.Phone,
		StoreDesc:     completeSeller.StoreDesc,
		Commission:    completeSeller.Commission,
		Address:       completeSeller.Address,
	}

	if err := Config.DB.Save(&sellerProfile).Error; err != nil {
		http.Error(w, "Error updating seller profile", http.StatusInternalServerError)
		return
	}
	//update reguser role field
	var user models.RegUser
	if err := Config.DB.First(&user, userID).Error; err == nil {
		if !strings.Contains(user.Role, "seller") {
			if user.Role != "" {
				user.Role += ",seller"
			} else {
				user.Role = "seller"
			}

		}
		Config.DB.Save(&user)
	}

	util.SendData(w, map[string]interface{}{
		"message": " seller profile updated sucessfully",
		"user_id": seller.ID,
	}, http.StatusAccepted)

}
