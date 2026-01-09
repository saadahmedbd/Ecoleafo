package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := util.RevokeRefreshToken(Config.DB, req.RefreshToken); err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Logout successfully", 200)
}
