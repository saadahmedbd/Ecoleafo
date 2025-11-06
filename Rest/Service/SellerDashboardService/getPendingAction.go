package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetPendingActions - Get items requiring attention
func (s *DashboardService) GetPendingActions(sellerID uint) (*sellerdashboard.PendingActionsResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	actions, err := s.DashboardRepository.GetPendingActions(sellerID)
	if err != nil {
		return nil, errors.New("failed to fetch pending actions")
	}

	return actions, nil
}
