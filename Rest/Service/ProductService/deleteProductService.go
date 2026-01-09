package productservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (s *ProductService) DeleteProduct(productID uint, userIdFromJWT uint, isAdmin bool) error {
	// Start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	fmt.Printf("DEBUG Service: productID=%d, userIdFromJWT=%d, isAdmin=%t\n", productID, userIdFromJWT, isAdmin)

	var product models.Product

	if isAdmin {
		// Admin can delete any product
		if err := s.db.Where("id = ?", productID).First(&product).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("product with ID %d not found", productID)
			}
			return fmt.Errorf("failed to find product: %v", err)
		}
	} else {
		// Regular seller - need to verify ownership
		// First, find the seller by user_id to get their actual ID
		var seller models.User
		if err := s.db.Where("user_id = ?", userIdFromJWT).First(&seller).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("seller with user_id %d does not exist", userIdFromJWT)
			}
			return fmt.Errorf("failed to verify seller: %v", err)
		}

		fmt.Printf("DEBUG Service: Found seller - ID=%d, user_id=%d, store=%s\n", seller.ID, seller.UserId, seller.StoreName)

		// Find the product and verify ownership using seller.ID
		if err := s.db.Where("id = ? AND seller_id = ?", productID, seller.ID).First(&product).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("product not found or you don't have permission to delete it")
			}
			return fmt.Errorf("failed to find product: %v", err)
		}
	}

	fmt.Printf("DEBUG Service: Found product - ID=%d, seller_id=%d, name=%s\n", product.ID, product.SellerID, product.Name)

	// Check if product has any dependencies that prevent deletion
	// You might want to prevent deletion if product has:
	// - Active orders
	// - Items in carts
	// etc.

	// Check for active orders (optional business logic)
	var orderCount int64
	tx.Model(&models.OrderItem{}).Where("product_id = ?", productID).Count(&orderCount)
	if orderCount > 0 {
		// Instead of preventing deletion, you might want to just mark as inactive
		// return fmt.Errorf("cannot delete product with existing orders")

		// Or soft delete and mark as inactive
		product.IsActive = false
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to deactivate product: %v", err)
		}
	}

	// Delete related records first (to avoid foreign key constraints)

	// Delete product images
	if err := tx.Where("product_id = ?", productID).Delete(&models.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete product images: %v", err)
	}

	// Delete product attributes
	if err := tx.Where("product_id = ?", productID).Delete(&models.ProductAttribute{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete product attributes: %v", err)
	}

	// Delete from wishlist
	if err := tx.Where("product_id = ?", productID).Delete(&models.Wishlist{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to remove from wishlists: %v", err)
	}

	// Remove from cart items (if any)
	if err := tx.Where("product_id = ?", productID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to remove from carts: %v", err)
	}

	// Delete reviews (optional - you might want to keep them)
	if err := tx.Where("product_id = ?", productID).Delete(&models.Review{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete reviews: %v", err)
	}

	// Finally, delete the product (soft delete due to gorm.DeletedAt)
	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete product: %v", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	// fmt.Printf("DEBUG Service: Successfully deleted product %d\n", productID)
	return nil
}
