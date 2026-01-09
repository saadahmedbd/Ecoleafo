package reviewhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// GetSellerProductReviews - Get all reviews for seller's products
func (h *ReviewHandler) GetSellerProductReviews(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get pagination params
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	perPage, _ := strconv.Atoi(r.URL.Query().Get("per_page"))
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	productIDStr := r.URL.Query().Get("product_id")
	var productID uint
	if productIDStr != "" {
		pid, _ := strconv.ParseUint(productIDStr, 10, 32)
		productID = uint(pid)
	}

	reviews, err := h.reviewService.GetSellerProductReviews(userID, productID, page, perPage)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, reviews, "Reviews fetched successfully")
}
