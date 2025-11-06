package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetOrderStatusDistribution - Get order distribution by status
func (s *DashboardService) GetOrderStatusDistribution(sellerID uint) ([]sellerdashboard.OrderStatusDistribution, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	distribution, err := s.DashboardRepository.GetOrderStatusDistribution(sellerID)
	if err != nil {
		return nil, errors.New("failed to fetch order distribution")
	}

	return distribution, nil
}
