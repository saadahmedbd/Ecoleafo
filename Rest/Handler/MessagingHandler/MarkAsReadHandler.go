package messaginghandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

// MarkAsRead - PUT /api/v1/messaging/conversations/{id}/read
func (h *MessagingHandler) MarkAsRead(w http.ResponseWriter, r *http.Request) {
	userCtx, err := h.getUserContext(r)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	convIDStr := r.URL.Query().Get("id")
	if convIDStr == "" {
		http.Error(w, "order id is required", http.StatusBadRequest)
		return
	}
	convID, err := strconv.Atoi(convIDStr)
	if err != nil {
		http.Error(w, "id not convert", http.StatusNoContent)
		return
	}

	if err := h.messageservice.MarkAsRead(userCtx, uint(convID)); err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Messages marked as read")
}
