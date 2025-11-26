package auditloghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs/recent?limit=20
func (h *AuditLogHandler) GetRecentActivity(w http.ResponseWriter, r *http.Request) {
	limit := util.ParseIntQuery(r, "limit", 20)

	logs, err := h.auditService.GetRecentActivity(limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get recent activity")
		return
	}

	util.RespondJSON(w, http.StatusOK, logs, "Recent activity retrieved")
}
