package sellerdashboardservice

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetLowStockProducts - Get products with low stock
func (s *DashboardService) GetLowStockProducts(sellerID uint, threshold int) ([]sellerdashboard.LowStockProductResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	if threshold <= 0 {
		threshold = 10 // Default threshold
	}

	products, err := s.DashboardRepository.GetLowStockProducts(sellerID, threshold)
	if err != nil {
		return nil, errors.New("failed to fetch low stock products")
	}

	return products, nil
}
