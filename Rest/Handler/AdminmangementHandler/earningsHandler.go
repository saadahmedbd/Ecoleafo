package adminmangementhandler

import (
	commissionpayoutearningservice "github.com/saadahmedbd/Treestore/Rest/Service/CommissionPayoutEarningService"
)

type AdminManagementHandler struct {
	commissionService commissionpayoutearningservice.CommissionService
}

func NewAdminManagementHandler(commissionService commissionpayoutearningservice.CommissionService) *AdminManagementHandler {
	return &AdminManagementHandler{
		commissionService: commissionService,
	}
}
