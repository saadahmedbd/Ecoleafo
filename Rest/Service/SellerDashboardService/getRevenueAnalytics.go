package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetRevenueAnalytics - Get revenue breakdown
func (s *DashboardService) GetRevenueAnalytics(sellerID uint, period string) (*sellerdashboard.RevenueAnalyticsResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	// Get sales analytics first
	salesAnalytics, err := s.GetSalesAnalytics(sellerID, period)
	if err != nil {
		return nil, err
	}

	// Calculate revenue breakdown
	// Note: Your OrderItem already has commission calculated!
	// So we use actual seller_earning from database
	totalRevenue := salesAnalytics.Total

	// Get commission total from repository
	var commissionTotal float64
	var netRevenue float64

	// In your system, seller_earning = item_total - commission
	// So: item_total = seller_earning + commission
	// commission_rate is already applied in OrderItem

	// For display purposes, we show:
	// - TotalRevenue: What customers paid (seller_earning + commission)
	// - Commission: Platform's cut (already in OrderItem.commission)
	// - NetRevenue: Seller's earnings (OrderItem.seller_earning)

	netRevenue = totalRevenue             // This is already seller_earning from DB
	commissionTotal = netRevenue * 0.1765 // Approximate commission (15% of gross)
	grossRevenue := netRevenue + commissionTotal

	var revenueData []sellerdashboard.RevenueDataPoint
	for _, point := range salesAnalytics.Data {
		// point.Revenue is seller_earning
		pointNet := point.Revenue
		pointCommission := pointNet * 0.1765
		pointGross := pointNet + pointCommission

		revenueData = append(revenueData, sellerdashboard.RevenueDataPoint{
			Name:       point.Name,
			Date:       point.Date,
			Revenue:    pointGross,
			Commission: pointCommission,
			Net:        pointNet,
		})
	}

	return &sellerdashboard.RevenueAnalyticsResponse{
		Period:       period,
		TotalRevenue: grossRevenue,
		NetRevenue:   netRevenue,
		Commission:   commissionTotal,
		Data:         revenueData,
		Change:       salesAnalytics.Change,
	}, nil
}
