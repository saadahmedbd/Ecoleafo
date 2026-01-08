package messaginghandler

import (
	"net/http"
	"strconv"
	"time"

	util "github.com/saadahmedbd/Treestore/Util"
)

// GetMessages - GET /api/v1/messaging/messages?conversation_id=X&since=timestamp&limit=50
func (h *MessagingHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	userCtx, err := h.getUserContext(r)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Parse query parameters
	convIDStr := r.URL.Query().Get("conversation_id")
	convID, err := strconv.ParseUint(convIDStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid conversation_id", http.StatusBadRequest)
		return
	}

	// Parse since timestamp (optional)
	var since time.Time
	sinceStr := r.URL.Query().Get("since")
	if sinceStr != "" {
		since, err = time.Parse(time.RFC3339, sinceStr)
		if err != nil {
			util.SendError(w, "Invalid since timestamp format (use RFC3339)", http.StatusBadRequest)
			return
		}
	}

	// Parse limit (optional, default 50)
	limit := 50
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err == nil && limitInt > 0 && limitInt <= 100 {
			limit = limitInt
		}
	}

	resp, err := h.messageservice.GetMessages(userCtx, uint(convID), since, limit)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, resp, "Messages retrieved successfully")
}
