package adminmangementhandler

import (
	"net/http"

	adminmangementservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminMangementService"
	util "github.com/saadahmedbd/Treestore/Util"
)

type DashboardHandler struct {
	dashboardService *adminmangementservice.DashboardService
}

func NewDashboardHandler(dashboardService *adminmangementservice.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *DashboardHandler) GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.dashboardService.GetDashboardStats()
	if err != nil {
		util.SendError(w, "Failed to fetch dashboard statistics", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Dashboard statistics retrieved successfully")
}
