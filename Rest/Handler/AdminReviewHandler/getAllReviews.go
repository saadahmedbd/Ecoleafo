package adminreviewhandler

import (
	"net/http"

	adminreviewdto "github.com/saadahmedbd/Treestore/Rest/DTO/AdminReviewDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/reviews?page=1&limit=20&status=pending&rating=1
func (h *AdminReviewHandler) GetAllReviews(w http.ResponseWriter, r *http.Request) {
	query := adminreviewdto.ReviewListQuery{
		Page:   util.ParseIntQuery(r, "page", 1),
		Limit:  util.ParseIntQuery(r, "limit", 20),
		Status: r.URL.Query().Get("status"),
		Rating: util.ParseIntQuery(r, "rating", 0),
		Search: r.URL.Query().Get("search"),
	}

	if r.URL.Query().Get("is_reported") == "true" {
		query.IsReported = true
	}

	reviews, total, err := h.adminreviewservice.GetAllReviews(query)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get reviews")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Reviews retrieved successfully", reviews, query.Page, query.Limit, int64(total))
}
