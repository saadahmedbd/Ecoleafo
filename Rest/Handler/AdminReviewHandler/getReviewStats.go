package adminreviewhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/reviews/stats
func (h *AdminReviewHandler) GetReviewStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.adminreviewservice.GetReviewStats()
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get review statistics")
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Review statistics retrieved successfully")
}
