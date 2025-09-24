package reviewhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (h *Handler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	reviewId := r.PathValue("reviewId")
	id, err := strconv.Atoi(reviewId)
	if err != nil {
		http.Error(w, "Invalid review id", http.StatusBadRequest)
		return
	}
	//try to delete
	var reviews models.Review
	if err := Config.DB.First(&reviews, id).Error; err != nil {
		http.Error(w, "order not found", http.StatusBadRequest)
		return
	}
	if err := Config.DB.Delete(&reviews).Error; err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}
	// util.SendData(w, seller, 200)
	w.Write([]byte(fmt.Sprintf("review %d deleted successfully", id)))

}
