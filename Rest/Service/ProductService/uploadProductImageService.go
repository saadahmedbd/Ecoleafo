package productservice

import (
	"fmt"
	"mime/multipart"

	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *ProductService) UploadProductImage(productID, userIdFromJWT uint, file multipart.File, fileHeader *multipart.FileHeader, altText string, isPrimary bool, sortOrder int) (*models.ProductImage, error) {
	// Verify product ownership
	if err := s.verifyProductOwnership(productID, userIdFromJWT); err != nil {
		return nil, err
	}

	// Validate file
	if err := s.validateImageFile(fileHeader); err != nil {
		return nil, err
	}

	// Upload to Cloudinary
	imageURL, err := util.UploadProductImage(file, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image: %v", err)
	}

	// If this is set as primary, unset other primary images
	if isPrimary {
		s.db.Model(&models.ProductImage{}).
			Where("product_id = ?", productID).
			Update("is_primary", false)
	}

	// Create image record in database
	productImage := models.ProductImage{
		ProductID: productID,
		ImageURL:  imageURL,
		AltText:   altText,
		IsPrimary: isPrimary,
		SortOrder: sortOrder,
	}

	if err := s.db.Create(&productImage).Error; err != nil {
		return nil, fmt.Errorf("failed to save image record: %v", err)
	}

	return &productImage, nil
}
