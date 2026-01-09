package auditloghandler

import (
	"net/http"
	"strconv"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// GET /api/admin/logs?page=1&limit=20&action=login&actor_type=admin
func (h *AuditLogHandler) GetLogs(w http.ResponseWriter, r *http.Request) {
	query := auditlogdto.AuditLogQuery{
		Page:        util.ParseIntQuery(r, "page", 1),
		Limit:       util.ParseIntQuery(r, "limit", 20),
		ActorType:   r.URL.Query().Get("actor_type"),
		Action:      r.URL.Query().Get("action"),
		ActionGroup: r.URL.Query().Get("action_group"),
		EntityType:  r.URL.Query().Get("entity_type"),
		Status:      r.URL.Query().Get("status"),
		Severity:    r.URL.Query().Get("severity"),
		Category:    r.URL.Query().Get("category"),
		StartDate:   r.URL.Query().Get("start_date"),
		EndDate:     r.URL.Query().Get("end_date"),
		Search:      r.URL.Query().Get("search"),
		IPAddress:   r.URL.Query().Get("ip"),
	}

	if actorID := r.URL.Query().Get("actor_id"); actorID != "" {
		id, _ := strconv.ParseUint(actorID, 10, 64)
		uid := uint(id)
		query.ActorID = &uid
	}
	if entityID := r.URL.Query().Get("entity_id"); entityID != "" {
		id, _ := strconv.ParseUint(entityID, 10, 64)
		uid := uint(id)
		query.EntityID = &uid
	}

	logs, total, err := h.auditService.GetLogs(query)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to get logs")
		return
	}

	util.SendPaginatedresponse(w, http.StatusOK, "Audit logs retrieved", logs, query.Page, query.Limit, int64(total))
}
