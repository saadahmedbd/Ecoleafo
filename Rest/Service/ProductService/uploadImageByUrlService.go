package productservice

import (
	"fmt"
	"strings"

	models "github.com/saadahmedbd/Treestore/Models"
)

type AddImageByURLRequest struct {
	ImageURL  string `json:"image_url" validate:"required,url"`
	AltText   string `json:"alt_text"`
	IsPrimary bool   `json:"is_primary"`
	SortOrder int    `json:"sort_order"`
}

func (s *ProductService) AddProductImageByURL(productID, userIdFromJWT uint, req AddImageByURLRequest) (*models.ProductImage, error) {
	// Verify product ownership
	if err := s.verifyProductOwnership(productID, userIdFromJWT); err != nil {
		return nil, err
	}

	// Validate URL format (basic validation)
	if !strings.HasPrefix(req.ImageURL, "http://") && !strings.HasPrefix(req.ImageURL, "https://") {
		return nil, fmt.Errorf("invalid image URL format")
	}

	// If this is set as primary, unset other primary images
	if req.IsPrimary {
		s.db.Model(&models.ProductImage{}).
			Where("product_id = ?", productID).
			Update("is_primary", false)
	}

	// Get sort order if not provided
	sortOrder := req.SortOrder
	if sortOrder <= 0 {

		s.db.Model(&models.ProductImage{}).
			Where("product_id = ?", productID).
			Select("COALESCE(MAX(sort_order), 0) + 1").
			Scan(&sortOrder)
	}

	// Create image record
	productImage := models.ProductImage{
		ProductID: productID,
		ImageURL:  req.ImageURL,
		AltText:   req.AltText,
		IsPrimary: req.IsPrimary,
		SortOrder: sortOrder,
	}

	if err := s.db.Create(&productImage).Error; err != nil {
		return nil, fmt.Errorf("failed to save image record: %v", err)
	}

	return &productImage, nil
}
