package adminreviewhandler

import (
	"encoding/json"
	"net/http"

	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// POST /api/reviews/report?id=5
func (h *AdminReviewHandler) ReportReview(w http.ResponseWriter, r *http.Request) {
	reviewID := util.ParseUintQuery(r, "id")
	if reviewID == 0 {
		util.RespondError(w, http.StatusBadRequest, "review id is required")
		return
	}

	var req adminreviewdto.ReportReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Get user data from context
	ctxUserID := r.Context().Value(constants.ContextKeyUserID)
	ctxRoles := r.Context().Value(constants.ContextKeyRole)

	if ctxUserID == nil || ctxRoles == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Convert roles
	roleList, ok := ctxRoles.([]interface{})
	if !ok {
		http.Error(w, "Invalid role type", http.StatusInternalServerError)
		return
	}

	// Check if buyer
	isBuyer := false
	for _, raw := range roleList {
		if r, ok := raw.(string); ok && r == "buyer" {
			isBuyer = true
			break
		}
	}

	if !isBuyer {
		http.Error(w, "Only buyer can access this", http.StatusForbidden)
		return
	}

	// Convert user id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	// Convert reguser.id → buyer.id
	buyerID, err := getBuyerID(uid)
	if err != nil {
		http.Error(w, `{"error":"admin account not found"}`, http.StatusNotFound)
		return
	}
	if err := h.adminreviewservice.ReportReview(reviewID, &req, buyerID, "buyer"); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusCreated, nil, "Review reported successfully")
}
