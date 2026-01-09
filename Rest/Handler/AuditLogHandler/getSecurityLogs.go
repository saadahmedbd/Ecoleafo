package auditloghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs/security?page=1&limit=20
func (h *AuditLogHandler) GetSecurityLogs(w http.ResponseWriter, r *http.Request) {
	page := util.ParseIntQuery(r, "page", 1)
	limit := util.ParseIntQuery(r, "limit", 20)

	logs, total, err := h.auditService.GetSecurityLogs(page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get security logs")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Security logs retrieved", logs, page, limit, int64(total))
}
