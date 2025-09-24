package cartitemhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetcartItemById(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.PathValue("cartitemId")
	sId, err := strconv.Atoi(cartItemId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var CartItems models.CartItem
	result := Config.DB.Preload("Products").First(&CartItems, sId)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	// not return password

	util.SendData(w, CartItems, 200)
}
