package buyerhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "please provide valid method", http.StatusMethodNotAllowed)
		return
	}
	var Buyers models.Buyer
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&Buyers)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	hashPassword, err := util.HashPassword(Buyers.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	Buyers.Password = hashPassword
	result := Config.DB.Create(&Buyers)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, Buyers, 200)
}
