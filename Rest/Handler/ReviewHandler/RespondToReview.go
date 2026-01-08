package reviewhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

type SellerResponseRequest struct {
	Response string `json:"response" validate:"required,min=10,max=1000"`
}

// RespondToReview - Seller responds to a review
func (h *ReviewHandler) RespondToReview(w http.ResponseWriter, r *http.Request) {
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

	reviewIDStr := r.PathValue("id")
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 32)
	if err != nil || reviewID == 0 {
		util.RespondError(w, http.StatusBadRequest, "invalid review ID")
		return
	}

	var req SellerResponseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "invalid request data")
		return
	}

	if err := util.ValidateStruct(req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	review, err := h.reviewService.RespondToReview(uint(reviewID), userID, req.Response)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, review, "Response added successfully")
}
