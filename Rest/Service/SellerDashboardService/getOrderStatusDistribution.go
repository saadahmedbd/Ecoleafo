package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetOrderStatusDistribution - Get order distribution by status
func (s *DashboardService) GetOrderStatusDistribution(regUserID uint) ([]sellerdashboard.OrderStatusDistribution, error) {
	if regUserID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	distribution, err := s.DashboardRepository.GetOrderStatusDistribution(regUserID)
	if err != nil {
		return nil, errors.New("failed to fetch order distribution")
	}

	return distribution, nil
}
