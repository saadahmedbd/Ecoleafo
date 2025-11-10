package reviewhandler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// DeleteReview - Delete a review

// DeleteReview - Delete a review
func (h *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	reviewIDStr := vars["id"]
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 32)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "invalid review ID")
		return
	}

	if err := h.reviewService.DeleteReview(uint(reviewID), regUserID); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "review deleted successfully")

}
