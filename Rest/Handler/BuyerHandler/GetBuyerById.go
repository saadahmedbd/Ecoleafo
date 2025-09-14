package buyerhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetBuyerById(w http.ResponseWriter, r *http.Request) {
	BuyerId := r.PathValue("buyerId")
	id, err := strconv.Atoi(BuyerId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var Buyers models.Buyer
	result := Config.DB.First(&Buyers, id)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	// not return password
	Buyers.Password = ""
	util.SendData(w, Buyers, 200)
}
