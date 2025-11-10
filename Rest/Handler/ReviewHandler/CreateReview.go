package reviewhandler

import (
	"encoding/json"
	"net/http"

	reviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/ReviewDTO"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// CreateReview - Create a new review
func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	// Get buyer ID from context (set by auth middleware)
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

	var req reviewdto.CreateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate request
	if err := util.ValidateStruct(req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	review, err := h.reviewService.CreateReview(regUserID, req)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, review, "Review Created successfully")

}
