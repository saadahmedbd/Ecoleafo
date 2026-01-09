package auditloghandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// POST /api/admin/logs/export
func (h *AuditLogHandler) ExportLogs(w http.ResponseWriter, r *http.Request) {
	var req auditlogdto.ExportAuditLogsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	data, filename, err := h.auditService.ExportLogs(req)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to export logs")
		return
	}

	contentType := "text/csv"
	if req.Format == "json" {
		contentType = "application/json"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Write(data)
}
