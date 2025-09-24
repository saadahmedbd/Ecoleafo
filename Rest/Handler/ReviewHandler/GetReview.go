package reviewhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid Method", http.StatusBadRequest)
		return
	}
	var reviews []models.Review

	err := Config.DB.Preload("Product").Preload("Buyer").Find(&reviews).Error
	if err != nil {
		http.Error(w, "Failed fetch review", http.StatusInternalServerError)
		return
	}
	util.SendData(w, reviews, 200)

}
