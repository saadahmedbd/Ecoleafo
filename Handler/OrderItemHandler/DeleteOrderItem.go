package orderitemhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	orderItemId := r.PathValue("orderitemId")
	id, err := strconv.Atoi(orderItemId)
	if err != nil {
		http.Error(w, "Invalid order item id", http.StatusBadRequest)
		return
	}
	//try to delete
	var orderItems models.OrderItem
	if err := Config.DB.First(&orderItems, id).Error; err != nil {
		http.Error(w, "order item not found", http.StatusBadRequest)
		return
	}
	if err := Config.DB.Delete(&orderItems).Error; err != nil {
		http.Error(w, "Failed to delete order item", http.StatusInternalServerError)
		return
	}
	// util.SendData(w, seller, 200)
	w.Write([]byte(fmt.Sprintf("order item %d deleted successfully", id)))

}
