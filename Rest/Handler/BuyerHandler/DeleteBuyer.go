package buyerhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (h *Handler) DeleteBuyer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	buyerId := r.PathValue("buyerId")
	id, err := strconv.Atoi(buyerId)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	//try to delete
	var buyers models.Buyer
	if err := Config.DB.First(&buyers, id).Error; err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	if err := Config.DB.Delete(&buyers).Error; err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	// util.SendData(w, seller, 200)
	w.Write([]byte(fmt.Sprintf("Buyer %d deleted successfully", id)))

}
