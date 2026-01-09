package sellerdashboardhandler

import sellerdashboardservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerDashboardService"

type DashboardHandler struct {
	DashboardService *sellerdashboardservice.DashboardService
}

func NewDashboardService(dashboardService *sellerdashboardservice.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		DashboardService: dashboardService,
	}
}
