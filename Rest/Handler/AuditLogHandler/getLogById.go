package auditloghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs/:id
func (h *AuditLogHandler) GetLogByID(w http.ResponseWriter, r *http.Request) {
	logID := util.ParseUintQuery(r, "id")
	if logID == 0 {
		util.RespondError(w, http.StatusBadRequest, "log id is required")
		return
	}

	log, err := h.auditService.GetLogByID(logID)
	if err != nil {
		util.RespondError(w, http.StatusNotFound, "Log not found")
		return
	}

	util.RespondJSON(w, http.StatusOK, log, "Audit log retrieved")
}
