package productservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type ProductFilters struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Status   string `json:"status"`   // active, inactive, approved, pending
	Category string `json:"category"` // category ID
	Search   string `json:"search"`   // search term
	SortBy   string `json:"sort_by"`  // name, price, created_at, updated_at
	Order    string `json:"order"`    // asc, desc
}

type ProductListResponse struct {
	Products   []models.Product `json:"products"`
	Pagination PaginationInfo   `json:"pagination"`
}

type PaginationInfo struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// GetSellerProducts - Get products for a specific seller
func (s *ProductService) GetSellerProducts(userIdFromJWT uint, filters ProductFilters) (*ProductListResponse, error) {
	// First, find the seller by user_id to get their actual ID
	var seller models.User
	if err := s.db.Where("user_id = ?", userIdFromJWT).First(&seller).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller with user_id %d does not exist", userIdFromJWT)
		}
		return nil, fmt.Errorf("failed to verify seller: %v", err)
	}

	// fmt.Printf("DEBUG Service: Found seller - ID=%d, user_id=%d, store=%s\n", seller.ID, seller.UserId, seller.StoreName)

	// Get products for this seller using their actual ID
	return s.getProductsBySellerID(seller.ID, filters, false)
}
