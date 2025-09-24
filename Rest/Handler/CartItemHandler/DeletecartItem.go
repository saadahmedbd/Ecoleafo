package cartitemhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (h *Handler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	cartItemId := r.PathValue("cartitemId")
	id, err := strconv.Atoi(cartItemId)
	if err != nil {
		http.Error(w, "Invalid cart item id", http.StatusBadRequest)
		return
	}
	//try to delete
	var cartitem models.CartItem
	if err := Config.DB.First(&cartitem, id).Error; err != nil {
		http.Error(w, "cart item not found", http.StatusBadRequest)
		return
	}
	if err := Config.DB.Delete(&cartitem).Error; err != nil {
		http.Error(w, "Failed to delete cart item", http.StatusInternalServerError)
		return
	}
	// util.SendData(w, seller, 200)
	w.Write([]byte(fmt.Sprintf("cartitem %d deleted successfully", id)))

}
