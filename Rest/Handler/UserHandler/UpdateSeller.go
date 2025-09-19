package userhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func UpdateSeller(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please provide valid method", http.StatusBadRequest)
		return
	}
	sellerId := r.PathValue("sellerId")
	id, err := strconv.Atoi(sellerId)
	fmt.Println("Params:", id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var existingSeller models.User
	if err := Config.DB.First(&existingSeller, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var updateData models.User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// If password is provided, hash it before updating
	if updateData.Password != "" {
		hashed, err := util.HashPassword(updateData.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		updateData.Password = hashed
	} else {
		updateData.Password = existingSeller.Password // keep old password
	}

	// Update existing seller with new values
	existingSeller.FirstName = updateData.FirstName
	existingSeller.LastName = updateData.LastName
	existingSeller.Email = updateData.Email
	existingSeller.Password = updateData.Password
	existingSeller.Phone = updateData.Phone
	existingSeller.StoreName = updateData.StoreName
	existingSeller.StoreDesc = updateData.StoreDesc
	existingSeller.IsActive = updateData.IsActive
	existingSeller.IsVerified = updateData.IsVerified

	if err := Config.DB.Save(&existingSeller).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Don’t expose password
	existingSeller.Password = ""
	util.SendData(w, updateData, 200)

}
