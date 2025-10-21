package authhandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, map[string]string{"message": "logged out successfully"}, 200)
}
