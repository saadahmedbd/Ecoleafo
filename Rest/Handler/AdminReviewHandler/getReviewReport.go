package adminreviewhandler

import (
	"fmt"
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/reviews/reports?status=pending&page=1&limit=20
func (h *AdminReviewHandler) GetReviewReports(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	page := util.ParseIntQuery(r, "page", 1)
	limit := util.ParseIntQuery(r, "limit", 20)

	reports, total, err := h.adminreviewservice.GetReviewReports(status, page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get review reports")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Review reports retrieved successfully", reports, page, limit, int64(total))
}

// POST /api/admin/reviews/reports/action?id=5&action=reviewed
func (h *AdminReviewHandler) ReviewReportAction(w http.ResponseWriter, r *http.Request) {
	reportID := util.ParseUintQuery(r, "id")
	action := r.URL.Query().Get("action")

	if reportID == 0 {
		util.RespondError(w, http.StatusBadRequest, "report id is required")
		return
	}

	if action != "reviewed" && action != "dismissed" {
		util.RespondError(w, http.StatusBadRequest, "action must be 'reviewed' or 'dismissed'")
		return
	}

	adminID := r.Context().Value("user_id").(uint)

	if err := h.adminreviewservice.ReviewReport(reportID, action, adminID); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, fmt.Sprintf("Report marked as %s", action))
}
