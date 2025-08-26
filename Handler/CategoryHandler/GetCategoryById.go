package categoryHandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("categoryId")
	sId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var categories models.Category
	result := Config.DB.Preload("Products").First(&categories, sId)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	// not return password

	util.SendData(w, categories, 200)
}
