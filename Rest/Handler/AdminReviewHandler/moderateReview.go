package adminreviewhandler

import (
	"encoding/json"
	"net/http"

	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/reviews/reported?page=1&limit=20

// POST /api/admin/reviews/moderate?id=5
func (h *AdminReviewHandler) ModerateReview(w http.ResponseWriter, r *http.Request) {
	reviewID := util.ParseUintQuery(r, "id")
	if reviewID == 0 {
		util.RespondError(w, http.StatusBadRequest, "review id is required")
		return
	}

	var req adminreviewdto.ModerateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	adminID := r.Context().Value("user_id").(uint)

	if err := h.adminreviewservice.ModerateReview(reviewID, &req, adminID); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	message := "Review approved successfully"
	if req.Action == "reject" {
		message = "Review rejected successfully"
	}

	util.RespondJSON(w, http.StatusOK, nil, message)
}
