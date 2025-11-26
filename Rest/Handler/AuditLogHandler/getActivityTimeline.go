package auditloghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs/timeline?start_date=2025-11-01&end_date=2025-11-30&group_by=day
func (h *AuditLogHandler) GetActivityTimeline(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	groupBy := r.URL.Query().Get("group_by")

	timeline, err := h.auditService.GetActivityTimeline(startDate, endDate, groupBy)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get timeline")
		return
	}

	util.RespondJSON(w, http.StatusOK, timeline, "Activity timeline retrieved")
}
