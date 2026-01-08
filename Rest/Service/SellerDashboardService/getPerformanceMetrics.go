package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetPerformanceMetrics - Get seller performance metrics
func (s *DashboardService) GetPerformanceMetrics(regUserID uint) (*sellerdashboard.PerformanceMetrics, error) {
	if regUserID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	metrics, err := s.DashboardRepository.GetPerformanceMetrics(regUserID)
	if err != nil {
		return nil, errors.New("failed to fetch performance metrics")
	}

	return metrics, nil
}
