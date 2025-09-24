package producthandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	sId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var products models.Product
	result := Config.DB.First(&products, sId)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	// not return password

	util.SendData(w, products, 200)
}
