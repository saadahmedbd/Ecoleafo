package sellerdashboardrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetPerformanceMetrics - Get performance indicators
func (r *DashboardRepository) GetPerformanceMetrics(regUserID uint) (*sellerdashboard.PerformanceMetrics, error) {
	// Convert reguser ID to seller ID
	var seller models.User
	if err := r.db.Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return nil, err
	}
	sellerID := seller.ID

	var metrics sellerdashboard.PerformanceMetrics

	// Get total orders for this seller
	var totalOrders int64
	r.db.Table("order_items").
		Select("COUNT(DISTINCT order_id)").
		Where("seller_id = ?", sellerID).
		Scan(&totalOrders)

	// Get completed/delivered orders
	var completedOrders int64
	r.db.Table("order_items").
		Select("COUNT(DISTINCT order_items.order_id)").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.status IN ?", sellerID, []string{"delivered"}).
		Scan(&completedOrders)

	// Calculate success rate
	if totalOrders > 0 {
		metrics.SuccessRate = (float64(completedOrders) / float64(totalOrders)) * 100
	}

	// Get average rating (from seller table already fetched)
	metrics.CustomerSatisfaction = seller.AverageRating

	// Calculate average order value (seller earnings)
	var totalEarnings float64
	r.db.Table("order_items").
		Select("COALESCE(SUM(seller_earning), 0)").
		Where("seller_id = ?", sellerID).
		Scan(&totalEarnings)

	if totalOrders > 0 {
		metrics.AverageOrderValue = totalEarnings / float64(totalOrders)
	}

	// Calculate fulfillment rate (orders shipped/delivered vs total)
	var fulfilledOrders int64
	r.db.Table("order_items").
		Select("COUNT(DISTINCT order_items.order_id)").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.status IN ?", sellerID, []string{"shipped", "delivered"}).
		Scan(&fulfilledOrders)

	if totalOrders > 0 {
		metrics.FulfillmentRate = (float64(fulfilledOrders) / float64(totalOrders)) * 100
	}

	// Calculate return rate (cancelled orders)
	var cancelledOrders int64
	r.db.Table("order_items").
		Select("COUNT(DISTINCT order_items.order_id)").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.status = ?", sellerID, "cancelled").
		Scan(&cancelledOrders)

	if totalOrders > 0 {
		metrics.ReturnRate = (float64(cancelledOrders) / float64(totalOrders)) * 100
	}

	// Conversion rate would require visitor data (placeholder)
	metrics.ConversionRate = 2.5 // Placeholder

	return &metrics, nil
}
