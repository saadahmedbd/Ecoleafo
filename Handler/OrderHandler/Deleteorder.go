package orderhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	orderId := r.PathValue("orderId")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}
	//try to delete
	var orders models.Order
	if err := Config.DB.First(&orders, id).Error; err != nil {
		http.Error(w, "order not found", http.StatusBadRequest)
		return
	}
	if err := Config.DB.Delete(&orders).Error; err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}
	// util.SendData(w, seller, 200)
	w.Write([]byte(fmt.Sprintf("order %d deleted successfully", id)))

}
