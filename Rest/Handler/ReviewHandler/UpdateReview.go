package reviewhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please provide valid method", http.StatusBadRequest)
		return
	}
	updateReviewId := r.PathValue("reviewId")
	id, err := strconv.Atoi(updateReviewId)

	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}
	var existingReview models.Review
	if err := Config.DB.First(&existingReview, id).Error; err != nil {
		http.Error(w, "review not found", http.StatusNotFound)
		return
	}
	var UpdateReview models.Review
	if err := json.NewDecoder(r.Body).Decode(&UpdateReview); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Update existing seller with new values
	existingReview.ProductID = UpdateReview.ProductID
	existingReview.BuyerID = UpdateReview.BuyerID
	existingReview.Rating = UpdateReview.Rating
	existingReview.Comment = UpdateReview.Comment

	if err := Config.DB.Save(&existingReview).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}
	util.SendData(w, UpdateReview, 200)
}
