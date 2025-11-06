package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetRecentOrders - Get recent orders with limit
func (s *DashboardService) GetRecentOrders(sellerID uint, limit int) ([]sellerdashboard.RecentOrderResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	if limit <= 0 || limit > 50 {
		limit = 5 // Default limit
	}

	orders, err := s.DashboardRepository.GetRecentOrders(sellerID, limit)
	if err != nil {
		return nil, errors.New("failed to fetch recent orders")
	}

	return orders, nil
}
