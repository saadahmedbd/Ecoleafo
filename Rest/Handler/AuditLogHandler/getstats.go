package auditloghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs/stats
func (h *AuditLogHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.auditService.GetStats()
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get statistics")
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Audit log statistics retrieved")
}
