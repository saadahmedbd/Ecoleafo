package messaginghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetUnreadCount - GET /api/v1/messaging/unread-count
func (h *MessagingHandler) GetUnreadCount(w http.ResponseWriter, r *http.Request) {
	userCtx, err := h.getUserContext(r)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp, err := h.messageservice.GetUnreadCount(userCtx)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, resp, "Unread count retrieved successfully")
}
