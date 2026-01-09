package adminmangementhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *AdminManagementHandler) GetTopEarningSellers(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitStr)

	if limit < 1 || limit > 100 {
		limit = 10
	}

	topSellers, err := h.commissionService.GetTopSellers(limit)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, topSellers, http.StatusOK)
}
