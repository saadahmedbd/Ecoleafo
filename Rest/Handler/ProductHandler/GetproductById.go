package producthandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"gorm.io/gorm"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	sId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	product, err := h.service.GetProductWithRelations(uint(sId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Id not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	go h.service.IncrementViewCount(uint(sId))
	// not return password

	util.SendData(w, product, 200)
}
