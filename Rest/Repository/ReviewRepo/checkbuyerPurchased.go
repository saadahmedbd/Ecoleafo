package reviewrepo

import (
	"fmt"
	"log"

	"github.com/saadahmedbd/Treestore/constants"
)

// CheckBuyerPurchased - Check if buyer purchased the product
// Returns: hasPurchased, buyerID, orderID, error
func (r *ReviewRepository) CheckBuyerPurchased(userID, productID uint) (bool, *uint, *uint, error) {
	// First, get the buyer_id from user_id
	var buyer struct {
		ID uint
	}
	err := r.db.Table("buyers").
		Select("id").
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		First(&buyer).Error

	if err != nil {
		log.Printf("DEBUG: No buyer found for user_id=%d: %v", userID, err)
		return false, nil, nil, fmt.Errorf("buyer account not found")
	}

	buyerID := buyer.ID
	log.Printf("DEBUG: user_id=%d maps to buyer_id=%d", userID, buyerID)

	// Check if buyer purchased and received the product
	var orderID uint
	err = r.db.Table("order_items").
		Select("order_items.order_id").
		Joins("INNER JOIN orders ON orders.id = order_items.order_id").
		Where("orders.buyer_id = ?", buyerID).
		Where("order_items.product_id = ?", productID).
		Where("orders.deleted_at IS NULL").
		Where("orders.status = ? OR order_items.status = ?", constants.OrderStatusDelivered, constants.OrderStatusDelivered).
		Limit(1).
		Pluck("order_items.order_id", &orderID).Error

	if err != nil {
		log.Printf("DEBUG: Error checking delivered status: %v", err)
		return false, nil, nil, err
	}

	if orderID == 0 {
		return false, nil, nil, fmt.Errorf("you have not purchased this product or it has not been delivered yet")
	}

	log.Printf("DEBUG: Order %d is delivered, allowing review", orderID)
	return true, &buyerID, &orderID, nil
}
