package cartitemhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetCartItemCount(w http.ResponseWriter, r *http.Request) {
	buyerId := r.PathValue("buyerId")
	id, err := strconv.Atoi(buyerId)
	if err != nil {
		http.Error(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}

	var count int64
	err = Config.DB.Model(&models.CartItem{}).Where("buyer_id = ?", id).Count(&count).Error
	if err != nil {
		http.Error(w, "Failed to get cart count", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"buyer_id":        id,
		"cart_item_count": count,
	}
	util.SendData(w, response, 200)
}