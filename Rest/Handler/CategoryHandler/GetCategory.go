package categoryHandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "please provide valid request", http.StatusBadRequest)
		return
	}
	var categories []models.Category
	err := Config.DB.Preload("Products").Find(&categories).Error
	if err != nil {
		http.Error(w, "can't find data", http.StatusBadGateway)
		return
	}
	util.SendData(w, categories, 200)
}
