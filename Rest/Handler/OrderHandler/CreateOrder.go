package orderhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	var orders models.Order
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&orders)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	result := Config.DB.Create(&orders)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, orders, 200)
}
