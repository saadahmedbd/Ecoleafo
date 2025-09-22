package authhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func Account(w http.ResponseWriter, r *http.Request) {
	var regAccount []models.RegUser
	if err := Config.DB.Find(&regAccount).Error; err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	util.SendData(w, regAccount, http.StatusOK)
}
