package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetSalesAnalytics - Get sales data for charts
func (s *DashboardService) GetSalesAnalytics(regUserID uint, period string) (*sellerdashboard.SalesAnalyticsResponse, error) {
	if regUserID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	// Validate period
	validPeriods := map[string]bool{
		"week":  true,
		"month": true,
		"year":  true,
	}

	if !validPeriods[period] {
		period = "week" // Default to week
	}

	analytics, err := s.DashboardRepository.GetSalesAnalytics(regUserID, period)
	if err != nil {
		return nil, errors.New("failed to fetch sales analytics")
	}

	return analytics, nil
}
