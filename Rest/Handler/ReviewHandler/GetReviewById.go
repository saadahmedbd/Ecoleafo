package reviewhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetReviewById(w http.ResponseWriter, r *http.Request) {
	reviewId := r.PathValue("reviewId")
	sId, err := strconv.Atoi(reviewId)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}
	var reviews models.Review
	result := Config.DB.Preload("product").Preload("Buyer").First(&reviews, sId)
	if result.Error != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	// not return password

	util.SendData(w, reviews, 200)
}
