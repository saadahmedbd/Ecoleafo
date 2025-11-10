package reviewrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// CheckBuyerPurchased - Check if buyer purchased the product
func (r *ReviewRepository) CheckBuyerPurchased(buyerID, productID uint) (bool, *uint, error) {
	var order models.Order
	err := r.db.Joins("JOIN order_items ON orders.id = order_items.order_id").
		Where("orders.buyer_id = ? AND order_items.product_id = ? AND orders.status = ?",
			buyerID, productID, "delivered").
		First(&order).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, &order.ID, nil
}
