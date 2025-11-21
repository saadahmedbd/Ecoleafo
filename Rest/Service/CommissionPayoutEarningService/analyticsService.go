package commissionpayoutearningservice

import (
	"time"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
)

// ============================================================================
func (s *commissionService) GetPlatformEarningsOverview() (*commissionearningpayoutdto.PlatformEarningsOverview, error) {
	data, err := s.commissionRepo.GetPlatformEarningsOverview()
	if err != nil {
		return nil, err
	}

	return &commissionearningpayoutdto.PlatformEarningsOverview{
		TotalGrossSales:       data["total_gross_sales"].(float64),
		TotalCommissionEarned: data["total_commission_earned"].(float64),
		TotalSellerEarnings:   data["total_seller_earnings"].(float64),
		PendingPayouts:        data["pending_payouts"].(float64),
		CompletedPayouts:      data["completed_payouts"].(float64),
		TotalOrders:           int(data["total_orders"].(int64)),
		CompletedOrders:       int(data["completed_orders"].(int64)),
	}, nil
}

func (s *commissionService) GetMonthlyRevenue(year int) ([]commissionearningpayoutdto.MonthlyRevenueReport, error) {
	if year == 0 {
		year = time.Now().Year()
	}

	data, err := s.commissionRepo.GetMonthlyRevenue(year)
	if err != nil {
		return nil, err
	}

	var response []commissionearningpayoutdto.MonthlyRevenueReport
	for _, item := range data {
		response = append(response, commissionearningpayoutdto.MonthlyRevenueReport{
			Month:            item["month"].(string),
			GrossRevenue:     item["gross_revenue"].(float64),
			CommissionEarned: item["commission_earned"].(float64),
			TotalOrders:      item["total_orders"].(int),
		})
	}

	return response, nil
}

func (s *commissionService) GetTopSellers(limit int) ([]commissionearningpayoutdto.TopSellerByRevenue, error) {
	if limit < 1 || limit > 100 {
		limit = 10
	}

	data, err := s.commissionRepo.GetTopSellersByRevenue(limit)
	if err != nil {
		return nil, err
	}

	var response []commissionearningpayoutdto.TopSellerByRevenue
	for _, item := range data {
		response = append(response, commissionearningpayoutdto.TopSellerByRevenue{
			SellerID:            item["seller_id"].(uint),
			SellerName:          item["seller_name"].(string),
			StoreName:           item["seller_name"].(string),
			TotalOrders:         item["total_orders"].(int),
			GrossSales:          item["gross_sales"].(float64),
			CommissionGenerated: item["commission_generated"].(float64),
		})
	}

	return response, nil
}
