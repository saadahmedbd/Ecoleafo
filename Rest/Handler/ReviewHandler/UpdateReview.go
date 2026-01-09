package reviewhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// UpdateReview - Update a review
func (h *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	userRoleVal := r.Context().Value(constants.ContextKeyRole)
	if userIDVal == nil || userRoleVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	regUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get review ID from URL
	reviewIDStr := r.PathValue("id")
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 32)
	if err != nil || reviewID == 0 {
		util.RespondError(w, http.StatusBadRequest, "invalid review ID")
		return
	}

	var req reviewdto.UpdateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "invalid request data")
		return
	}

	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "validation failed")
		return
	}

	review, err := h.reviewService.UpdateReview(uint(reviewID), regUserID, req)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, review, "Review updated successfully")

}
