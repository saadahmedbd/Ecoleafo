package sellerdashboardrepo

import (
	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
	util "github.com/saadahmedbd/Treestore/Util"
)

// GetRecentOrders - Get recent orders for dashboard
func (r *DashboardRepository) GetRecentOrders(sellerID uint, limit int) ([]sellerdashboard.RecentOrderResponse, error) {
	var orders []sellerdashboard.RecentOrderResponse

	err := r.db.Table("orders").
		Select(`DISTINCT orders.id, orders.order_number, 
			CONCAT(reg_users.first_name, ' ', reg_users.last_name) as customer_name,
			order_items.product_name, 
			order_items.seller_earning as total, 
			orders.status, 
			orders.created_at`).
		Joins("JOIN buyers ON orders.buyer_id = buyers.id").
		Joins("JOIN reg_users ON buyers.user_id = reg_users.id").
		Joins("JOIN order_items ON orders.id = order_items.order_id").
		Where("order_items.seller_id = ?", sellerID).
		Order("orders.created_at DESC").
		Limit(limit).
		Scan(&orders).Error

	if err != nil {
		return nil, err
	}

	// Calculate time ago for each order
	for i := range orders {
		orders[i].TimeAgo = util.FormatTimeAgo(orders[i].CreatedAt)
	}

	return orders, nil
}
