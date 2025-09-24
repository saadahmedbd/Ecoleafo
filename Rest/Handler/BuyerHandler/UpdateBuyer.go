package buyerhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) UpdateBuyer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "plaease provide valid method", http.StatusBadRequest)
		return
	}
	buyerId := r.PathValue("buyerId")
	id, err := strconv.Atoi(buyerId)
	if err != nil {
		http.Error(w, "cannot convert id to int", http.StatusBadRequest)
		return
	}
	var existingBuyer models.Buyer
	if err := Config.DB.First(&existingBuyer, id).Error; err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	// decode the json value
	var UpdateBuyers models.Buyer
	if err := json.NewDecoder(r.Body).Decode(&UpdateBuyers); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// If password is provided, hash it before updating
	if UpdateBuyers.Password != "" {
		hashed, err := util.HashPassword(UpdateBuyers.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		UpdateBuyers.Password = hashed
	} else {
		UpdateBuyers.Password = existingBuyer.Password // keep old password
	}

	existingBuyer.Password = UpdateBuyers.Password
	existingBuyer.Phone = UpdateBuyers.Phone
	existingBuyer.DefaultAddress = UpdateBuyers.DefaultAddress

	if err := Config.DB.Save(&existingBuyer).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Don’t expose password
	existingBuyer.Password = ""
	util.SendData(w, UpdateBuyers, 200)
}
