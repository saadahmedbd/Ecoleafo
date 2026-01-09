package userhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (h *Handler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	sellerId := r.PathValue("sellerId")
	id, err := strconv.Atoi(sellerId)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	//try to delete
	var seller models.User
	if err := Config.DB.First(&seller, id).Error; err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	if err := Config.DB.Delete(&seller).Error; err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	// util.SendData(w, seller, 200)
	w.Write([]byte(fmt.Sprintf("User %d deleted successfully", id)))

}
