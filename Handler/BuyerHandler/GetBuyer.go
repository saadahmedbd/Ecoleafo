package buyerhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetBuyer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid Method", http.StatusBadRequest)
		return
	}
	var buyers []models.Buyer

	err := Config.DB.Preload("Role").Preload("Orders").Preload("CartItems").Preload("Reviews").Find(&buyers).Error
	if err != nil {
		http.Error(w, "Failed fetch user", http.StatusInternalServerError)
		return
	}
	util.SendData(w, buyers, 200)
}
