package sellerdashboardrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetPendingActions - Get items needing attention
func (r *DashboardRepository) GetPendingActions(regUserID uint) (*sellerdashboard.PendingActionsResponse, error) {
	// Convert reguser ID to seller ID
	var seller struct {
		ID uint
	}
	if err := r.db.Table("users").Select("id").Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return nil, err
	}
	sellerID := seller.ID

	var actions sellerdashboard.PendingActionsResponse

	// Pending orders (orders with pending status containing seller's products)
	var pendingOrders int64
	r.db.Table("order_items").
		Select("COUNT(DISTINCT order_items.order_id)").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.status = ?", sellerID, "pending").
		Scan(&pendingOrders)
	actions.PendingOrders = int(pendingOrders)

	// Pending products (not approved)
	var pendingProducts int64
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND is_approved = ?", sellerID, false).
		Count(&pendingProducts)
	actions.PendingProducts = int(pendingProducts)

	// Out of stock
	var outOfStock int64
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND quantity = ?", sellerID, 0).
		Count(&outOfStock)
	actions.OutOfStock = int(outOfStock)

	// Low stock
	var lowStock int64
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND quantity < ? AND quantity > ?", sellerID, 10, 0).
		Count(&lowStock)
	actions.LowStock = int(lowStock)

	return &actions, nil
}
