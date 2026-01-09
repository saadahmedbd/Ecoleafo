package cartitemhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) CreatecartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	var CartItems models.CartItem
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&CartItems)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	result := Config.DB.Create(&CartItems)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, CartItems, 200)
}
