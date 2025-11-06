package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetTopProducts - Get best selling products
func (s *DashboardService) GetTopProducts(sellerID uint, limit int) ([]sellerdashboard.TopProductResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	if limit <= 0 || limit > 20 {
		limit = 5 // Default limit
	}

	products, err := s.DashboardRepository.GetTopProducts(sellerID, limit)
	if err != nil {
		return nil, errors.New("failed to fetch top products")
	}

	return products, nil
}
