package sellerdashboardrepo

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// UpdateOrderStatus - Update order status
func (r *DashboardRepository) UpdateOrderStatus(regUserID uint, orderID uint, status string, trackingNumber string, notes string) error {
	// Convert reguser ID to seller ID
	var seller struct {
		ID uint
	}
	if err := r.db.Table("users").Select("id").Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return err
	}
	sellerID := seller.ID

	// Verify this order belongs to this seller
	var count int64
	r.db.Table("order_items").
		Where("order_id = ? AND seller_id = ?", orderID, sellerID).
		Count(&count)

	if count == 0 {
		return errors.New("order not found or does not belong to this seller")
	}

	// Update order status
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	if status == "shipped" && trackingNumber != "" {
		updates["tracking_number"] = trackingNumber
		updates["shipped_at"] = time.Now()
	}

	if status == "delivered" {
		updates["delivered_at"] = time.Now()
	}

	if err := r.db.Model(&models.Order{}).Where("id = ?", orderID).Updates(updates).Error; err != nil {
		return err
	}

	// Create order history
	history := &models.OrderHistory{
		OrderID:   orderID,
		Status:    models.OrderStatusEnum(status),
		Comment:   notes,
		UpdatedBy: fmt.Sprintf("seller_%d", sellerID),
	}

	return r.db.Create(history).Error
}
