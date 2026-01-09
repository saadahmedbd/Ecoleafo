package auditloghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs/actor?actor_id=5&actor_type=admin&page=1&limit=20
func (h *AuditLogHandler) GetActorLogs(w http.ResponseWriter, r *http.Request) {
	actorID := util.ParseUintQuery(r, "actor_id")
	actorType := r.URL.Query().Get("actor_type")
	page := util.ParseIntQuery(r, "page", 1)
	limit := util.ParseIntQuery(r, "limit", 20)

	if actorID == 0 {
		util.RespondError(w, http.StatusBadRequest, "actor_id is required")
		return
	}

	logs, total, err := h.auditService.GetActorLogs(actorID, actorType, page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get actor logs")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Actor logs retrieved", logs, page, limit, int64(total))
}
