package adminreviewhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/reviews/reported?page=1&limit=20
func (h *AdminReviewHandler) GetReportedReviews(w http.ResponseWriter, r *http.Request) {
	page := util.ParseIntQuery(r, "page", 1)
	limit := util.ParseIntQuery(r, "limit", 20)

	reviews, total, err := h.adminreviewservice.GetReportedReviews(page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get reported reviews")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Reported reviews retrieved successfully", reviews, page, limit, int64(total))
}
