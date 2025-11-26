package auditloghandler

import (
	"fmt"
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// POST /api/admin/logs/cleanup?days=90
func (h *AuditLogHandler) CleanupOldLogs(w http.ResponseWriter, r *http.Request) {
	days := util.ParseIntQuery(r, "days", 90)

	if err := h.auditService.CleanupOldLogs(days); err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to cleanup logs")
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, fmt.Sprintf("Logs older than %d days have been deleted", days))
}
