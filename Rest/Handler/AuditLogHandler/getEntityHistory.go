package auditloghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs/entity?entity_type=product&entity_id=10&page=1&limit=20
func (h *AuditLogHandler) GetEntityHistory(w http.ResponseWriter, r *http.Request) {
	entityType := r.URL.Query().Get("entity_type")
	entityID := util.ParseUintQuery(r, "entity_id")
	page := util.ParseIntQuery(r, "page", 1)
	limit := util.ParseIntQuery(r, "limit", 20)

	if entityType == "" || entityID == 0 {
		util.RespondError(w, http.StatusBadRequest, "entity_type and entity_id are required")
		return
	}

	logs, total, err := h.auditService.GetEntityHistory(entityType, entityID, page, limit)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get entity history")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Entity history retrieved", logs, page, limit, int64(total))
}
