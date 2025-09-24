package rolehandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	var Roles []models.Role
	db := Config.DB
	db.Find(&Roles)
	util.SendData(w, Roles, 200)
}
