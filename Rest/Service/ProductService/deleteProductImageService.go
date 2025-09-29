package productservice

import (
	"errors"
	"fmt"
	"os"
	"strings"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (s *ProductService) DeleteProductImage(productID, imageID, userIdFromJWT uint) error {
	// Verify product ownership
	if err := s.verifyProductOwnership(productID, userIdFromJWT); err != nil {
		return err
	}

	// Find the image
	var productImage models.ProductImage
	if err := s.db.Where("id = ? AND product_id = ?", imageID, productID).First(&productImage).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("image not found")
		}
		return fmt.Errorf("failed to find image: %v", err)
	}

	// Delete file from disk (if it's a local upload)
	if strings.HasPrefix(productImage.ImageURL, "/uploads/") {
		filePath := "." + productImage.ImageURL // Adjust path as needed
		os.Remove(filePath)                     // Ignore error if file doesn't exist
	}

	// Delete from database
	if err := s.db.Delete(&productImage).Error; err != nil {
		return fmt.Errorf("failed to delete image record: %v", err)
	}

	return nil
}
