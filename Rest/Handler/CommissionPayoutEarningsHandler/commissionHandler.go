package commissionpayoutearningshandler

import commissionpayoutearningservice "github.com/saadahmedbd/Treestore/Rest/Service/CommissionPayoutEarningService"

type CommissionHandler struct {
	commissionService commissionpayoutearningservice.CommissionService
}

func NewCommissionHandler(commissionService commissionpayoutearningservice.CommissionService) *CommissionHandler {
	return &CommissionHandler{
		commissionService: commissionService,
	}
}
