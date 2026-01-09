package sellerdashboardservice

import sellerdashboardrepo "github.com/saadahmedbd/Treestore/Rest/Repository/SellerDashboardRepo"

type DashboardService struct {
	DashboardRepository *sellerdashboardrepo.DashboardRepository
}

func NewDashboardService(dashboardRepo *sellerdashboardrepo.DashboardRepository) *DashboardService {
	return &DashboardService{
		DashboardRepository: dashboardRepo,
	}
}
