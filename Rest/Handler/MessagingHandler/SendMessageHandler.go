package messaginghandler

import (
	"encoding/json"
	"net/http"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// SendMessage - POST /api/v1/messaging/messages
func (h *MessagingHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	userCtx, err := h.getUserContext(r)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req messagingdto.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		util.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.messageservice.SendMessage(userCtx, &req)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusCreated, resp, "Message sent successfully")
}
