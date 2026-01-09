package productservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (s *ProductService) GetPublicSellerProducts(sellerID uint, filters ProductFilters) (*ProductListResponse, error) {
	// Verify seller exists and is active
	// 	/api/sellers/{seller_id}/products - Anyone can view seller's products
	// Shows only ACTIVE and APPROVED products
	// No authentication required
	var seller models.User
	if err := s.db.Where("id = ? AND is_active = ? AND is_approved = ?", sellerID, true, true).First(&seller).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller not found or not active")
		}
		return nil, fmt.Errorf("failed to verify seller: %v", err)
	}
	// Only show active and approved products to public
	filters.Status = "active"
	return s.getProductsBySellerID(sellerID, filters, true)
}
