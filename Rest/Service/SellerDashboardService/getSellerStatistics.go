package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetSellerStatistics - Get comprehensive dashboard statistics
func (s *DashboardService) GetSellerStatistics(sellerID uint) (*sellerdashboard.DashboardStatsResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	stats, err := s.DashboardRepository.GetSellerStatistics(sellerID)
	if err != nil {
		return nil, errors.New("failed to fetch statistics")
	}

	return stats, nil
}
