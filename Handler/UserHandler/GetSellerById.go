package userhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetSellerById(w http.ResponseWriter, r *http.Request) {
	sellerId := r.PathValue("sellerId")
	sId, err := strconv.Atoi(sellerId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var Sellers models.User
	result := Config.DB.First(&Sellers, sId)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	// not return password
	Sellers.Password = ""
	util.SendData(w, Sellers, 200)
}
