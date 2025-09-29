package productservice

import (
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (s *ProductService) verifyProductOwnership(productID, userIdFromJWT uint) error {
	// Find seller by user_id
	var seller models.User
	if err := s.db.Where("user_id = ?", userIdFromJWT).First(&seller).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("seller not found")
		}
		return fmt.Errorf("failed to verify seller: %v", err)
	}

	// Verify product ownership
	var product models.Product
	if err := s.db.Where("id = ? AND seller_id = ?", productID, seller.ID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found or access denied")
		}
		return fmt.Errorf("failed to verify product ownership: %v", err)
	}

	return nil
}

func (s *ProductService) validateImageFile(fileHeader *multipart.FileHeader) error {
	// Check file size (max 5MB)
	maxSize := int64(5 << 20) // 5MB
	if fileHeader.Size > maxSize {
		return fmt.Errorf("file too large (max 5MB)")
	}

	// Check file extension
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExts[ext] {
		return fmt.Errorf("invalid file type (allowed: jpg, jpeg, png, gif, webp)")
	}

	return nil
}

func (s *ProductService) generateImageFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("product_%d_%s%s", timestamp, generateRandomString(8), ext)
}

func generateRandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
