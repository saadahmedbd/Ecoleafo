package messaginghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetConversations - GET /api/v1/messaging/conversations
func (h *MessagingHandler) GetConversations(w http.ResponseWriter, r *http.Request) {
	userCtx, err := h.getUserContext(r)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp, err := h.messageservice.GetConversations(userCtx)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, resp, "Conversations retrieved successfully")
}
