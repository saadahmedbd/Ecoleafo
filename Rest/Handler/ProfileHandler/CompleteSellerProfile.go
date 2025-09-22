package profilehandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

type ComPeleteSeller struct {
	UserId uint `json:"user_id"`
	// Add other profile fields as needed
	StoreName  string  `json:"business_name"`
	Address    string  `json:"address"`
	Phone      string  `json:"phone"`
	StoreDesc  string  `json:"store_description"`
	Commission float64 `json:"commission"`
}

func CompleteSeller(w http.ResponseWriter, r *http.Request) {
	var completeSeller ComPeleteSeller
	encode := json.NewDecoder(r.Body)
	err := encode.Decode(&completeSeller)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	// Check if seller profile already exists
	// If not, create a new seller profile
	var seller models.User
	if err := Config.DB.Where("user_id= ?", completeSeller.UserId).First(&seller).Error; err != nil {
		http.Error(w, "Seller profile not found", http.StatusNotFound)
		return
	}
	//update seller feild
	seller.StoreName = completeSeller.StoreName
	seller.Phone = completeSeller.Phone
	seller.StoreDesc = completeSeller.StoreDesc
	seller.Commission = completeSeller.Commission
	if err := Config.DB.Save(&seller).Error; err != nil {
		http.Error(w, "Error updating seller profile", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]interface{}{
		"message": " seller profile updated sucessfully",
		"user_id": seller.ID,
	}, http.StatusAccepted)

}
