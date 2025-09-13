package cartitemhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func GetCartItemsByBuyer(w http.ResponseWriter, r *http.Request) {
	buyerId := r.PathValue("buyerId")
	id, err := strconv.Atoi(buyerId)
	if err != nil {
		http.Error(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}

	var cartItems []models.CartItem
	err = Config.DB.Preload("Product").Where("buyer_id = ?", id).Find(&cartItems).Error
	if err != nil {
		http.Error(w, "Failed to fetch cart items", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"buyer_id":        id,
		"cart_items":      cartItems,
		"cart_item_count": len(cartItems),
	}
	util.SendData(w, response, 200)
}