package userhandler

import (
	"encoding/json"
	"net/http"

	config "github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func CreateSeller(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "please provide valid method", http.StatusMethodNotAllowed)
		return
	}
	var Sellers models.User
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&Sellers)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	hashPassword, err := util.HashPassword(Sellers.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	Sellers.Password = hashPassword
	result := config.DB.Create(&Sellers)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, Sellers, 200)

}
