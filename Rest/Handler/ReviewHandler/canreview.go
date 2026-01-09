package reviewhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// CanReview - Check if buyer can review product
func (h *ReviewHandler) CanReview(w http.ResponseWriter, r *http.Request) {
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

	productIDStr := r.URL.Query().Get("product_id")
	if productIDStr == "" {
		util.RespondError(w, http.StatusBadRequest, "product_id is required")
		return
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil || productID == 0 {
		util.RespondError(w, http.StatusBadRequest, "invalid product_id")
		return
	}

	canReview, err := h.reviewService.CanReview(regUserID, uint(productID))
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, canReview, "Review Created successfully")

}
