package reviewhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// PublicAPIResponse represents the public API response structure
type PublicAPIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// GetProductReviews - Get reviews for a product
func (h *ReviewHandler) GetProductReviews(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Parse product_id
	productIDStr := query.Get("product_id")
	if productIDStr == "" {
		util.RespondError(w, http.StatusBadRequest, "product_id is required")
		return
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "invalid product_id")
		return
	}

	// Parse optional filters
	filter := reviewdto.ReviewFilterRequest{
		ProductID: uint(productID),
		Page:      1,
		PerPage:   10,
		SortBy:    "newest",
	}

	if ratingStr := query.Get("rating"); ratingStr != "" {
		if rating, err := strconv.Atoi(ratingStr); err == nil {
			filter.Rating = rating
		}
	}

	if pageStr := query.Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			filter.Page = page
		}
	}

	if perPageStr := query.Get("per_page"); perPageStr != "" {
		if perPage, err := strconv.Atoi(perPageStr); err == nil && perPage > 0 && perPage <= 50 {
			filter.PerPage = perPage
		}
	}

	if sortBy := query.Get("sort_by"); sortBy != "" {
		filter.SortBy = sortBy
	}

	reviews, err := h.reviewService.GetReviewsByProduct(filter)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Send response in the expected format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := PublicAPIResponse{
		Success: true,
		Data:    reviews,
	}

	json.NewEncoder(w).Encode(response)
}
