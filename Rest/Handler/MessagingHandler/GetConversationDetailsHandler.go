package messaginghandler

import (
	util "github.com/saadahmedbd/Treestore/Util"
	"net/http"
	"strconv"
)

// GetConversationDetails - GET /api/v1/messaging/conversations/{id}
func (h *MessagingHandler) GetConversationDetails(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.messageservice.GetConversationDetails(userCtx, uint(convID))
	if err != nil {
		util.SendError(w, err.Error(), http.StatusNotFound)
		return
	}

	util.RespondJSON(w, http.StatusOK, resp, "Conversation details retrieved successfully")
}
