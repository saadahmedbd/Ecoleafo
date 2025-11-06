package sellerdashboardrepo

import sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"

// GetOrderStatusDistribution - Get order counts by status
func (r *DashboardRepository) GetOrderStatusDistribution(sellerID uint) ([]sellerdashboard.OrderStatusDistribution, error) {
	var distribution []sellerdashboard.OrderStatusDistribution

	type StatusCount struct {
		Status string
		Count  int64
	}
	var statusCounts []StatusCount

	err := r.db.Table("order_items").
		Select("orders.status, COUNT(DISTINCT order_items.order_id) as count").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ?", sellerID).
		Group("orders.status").
		Scan(&statusCounts).Error

	if err != nil {
		return nil, err
	}

	// Map to response with colors
	colorMap := map[string]string{
		"delivered":  "#10B981",
		"pending":    "#F59E0B",
		"processing": "#3B82F6",
		"shipped":    "#6366F1",
		"cancelled":  "#EF4444",
	}

	for _, sc := range statusCounts {
		distribution = append(distribution, sellerdashboard.OrderStatusDistribution{
			Status: sc.Status,
			Count:  int(sc.Count),
			Color:  colorMap[sc.Status],
			Value:  float64(sc.Count),
		})
	}

	return distribution, nil
}
