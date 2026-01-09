package productservice

import (
	"fmt"
	"mime/multipart"
	"strings"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (s *ProductService) UploadMultipleProductImages(productID, userIdFromJWT uint, files []*multipart.FileHeader) ([]*models.ProductImage, error) {
	// Verify product ownership
	if err := s.verifyProductOwnership(productID, userIdFromJWT); err != nil {
		return nil, err
	}

	var imageList []*models.ProductImage
	var uploadErrors []string

	// Get current max sort order
	var maxSortOrder int
	s.db.Model(&models.ProductImage{}).
		Where("product_id = ?", productID).
		Select("COALESCE(MAX(sort_order), 0)").
		Scan(&maxSortOrder)

	for i, fileHeader := range files {
		// Validate each file
		if err := s.validateImageFile(fileHeader); err != nil {
			uploadErrors = append(uploadErrors, fmt.Sprintf("File %s: %v", fileHeader.Filename, err))
			continue
		}

		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			uploadErrors = append(uploadErrors, fmt.Sprintf("File %s: failed to open", fileHeader.Filename))
			continue
		}

		// Upload single image
		altText := fmt.Sprintf("Product image %d", i+1)
		sortOrder := maxSortOrder + i + 1
		isPrimary := i == 0 && len(imageList) == 0 // First image is primary if no images exist

		imageData, err := s.UploadProductImage(productID, userIdFromJWT, file, fileHeader, altText, isPrimary, sortOrder)
		file.Close()

		if err != nil {
			uploadErrors = append(uploadErrors, fmt.Sprintf("File %s: %v", fileHeader.Filename, err))
			continue
		}

		imageList = append(imageList, imageData)
	}

	// Return results even if some uploads failed
	if len(uploadErrors) > 0 {
		return imageList, fmt.Errorf("some uploads failed: %s", strings.Join(uploadErrors, "; "))
	}

	return imageList, nil
}
