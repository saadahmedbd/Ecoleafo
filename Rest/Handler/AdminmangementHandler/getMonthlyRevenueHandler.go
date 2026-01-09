package adminmangementhandler

import (
	"net/http"
	"strconv"
	"time"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *AdminManagementHandler) GetMonthlyRevenue(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	year, _ := strconv.Atoi(yearStr)

	if year == 0 {
		year = time.Now().Year()
	}

	revenue, err := h.commissionService.GetMonthlyRevenue(year)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, revenue, http.StatusOK)
}
