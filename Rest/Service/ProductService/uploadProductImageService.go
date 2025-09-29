package productservice

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	models "github.com/saadahmedbd/Treestore/Models"
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

	// Generate unique filename
	fileName := s.generateImageFileName(fileHeader.Filename)

	// Create upload directory if it doesn't exist
	uploadDir := "uploads/products"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Save file to disk
	filePath := filepath.Join(uploadDir, fileName)
	destFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %v", err)
	}
	defer destFile.Close()

	// Copy file content
	_, err = io.Copy(destFile, file)
	if err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	// If this is set as primary, unset other primary images
	if isPrimary {
		s.db.Model(&models.ProductImage{}).
			Where("product_id = ?", productID).
			Update("is_primary", false)
	}

	// Create image record in database
	imageURL := fmt.Sprintf("/uploads/products/%s", fileName) // Adjust based on your URL structure
	productImage := models.ProductImage{
		ProductID: productID,
		ImageURL:  imageURL,
		AltText:   altText,
		IsPrimary: isPrimary,
		SortOrder: sortOrder,
	}

	if err := s.db.Create(&productImage).Error; err != nil {
		// Delete uploaded file if database insert fails
		os.Remove(filePath)
		return nil, fmt.Errorf("failed to save image record: %v", err)
	}

	return &productImage, nil
}
