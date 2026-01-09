package adminreviewhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// DELETE /api/admin/reviews/delete?id=5
func (h *AdminReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	reviewID := util.ParseUintQuery(r, "id")
	if reviewID == 0 {
		util.RespondError(w, http.StatusBadRequest, "review id is required")
		return
	}

	adminID := r.Context().Value("user_id").(uint)

	if err := h.adminreviewservice.DeleteReview(reviewID, adminID); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Review deleted successfully")
}
