package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetOrderDetails - Get order details by ID
func (s *DashboardService) GetOrderDetails(regUserID uint, orderID uint) (*sellerdashboard.OrderDetailsResponse, error) {
	if regUserID == 0 || orderID == 0 {
		return nil, errors.New("invalid parameters")
	}

	order, err := s.DashboardRepository.GetOrderDetails(regUserID, orderID)
	if err != nil {
		return nil, errors.New("failed to fetch order details")
	}

	return order, nil
}
