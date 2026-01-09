package messaginghandler

import (
	"encoding/json"
	"net/http"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// CreateConversation - POST /api/v1/messaging/conversations
func (h *MessagingHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	userCtx, err := h.getUserContext(r)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req messagingdto.CreateConversationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		util.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.messageservice.CreateConversation(userCtx, &req)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusCreated, resp, "Conversation created successfully")
}
