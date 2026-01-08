package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetPendingActions - Get items requiring attention
func (s *DashboardService) GetPendingActions(regUserID uint) (*sellerdashboard.PendingActionsResponse, error) {
	if regUserID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	actions, err := s.DashboardRepository.GetPendingActions(regUserID)
	if err != nil {
		return nil, errors.New("failed to fetch pending actions")
	}

	return actions, nil
}
