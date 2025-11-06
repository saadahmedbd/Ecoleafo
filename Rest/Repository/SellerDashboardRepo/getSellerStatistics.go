package sellerdashboardrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetSellerStatistics - Get comprehensive seller statistics
func (r *DashboardRepository) GetSellerStatistics(sellerID uint) (*sellerdashboard.DashboardStatsResponse, error) {
	var stats sellerdashboard.DashboardStatsResponse

	// Get seller info for averageRating and totalReviews
	var seller models.User
	if err := r.db.Where("user_id = ?", sellerID).First(&seller).Error; err != nil {
		return nil, err
	}

	stats.AverageRating = seller.AverageRating
	stats.TotalReviews = seller.TotalReviews
	stats.TotalSales = seller.TotalSales
	stats.TotalEarnings = seller.TotalEarnings
	stats.TotalOrders = seller.TotalOrders

	// Get order statistics from OrderItems (since orders are split by seller)
	var orderStats []struct {
		Status string
		Count  int64
	}

	// Count orders by status from order_items table
	r.db.Table("order_items").
		Select("orders.status, COUNT(DISTINCT order_items.order_id) as count").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ?", sellerID).
		Group("orders.status").
		Scan(&orderStats)

	for _, stat := range orderStats {
		switch stat.Status {
		case "pending":
			stats.PendingOrders = int(stat.Count)
		case "delivered":
			stats.CompletedOrders += int(stat.Count)
		case "shipped":
			stats.ShippedOrders = int(stat.Count)
		case "cancelled":
			stats.CancelledOrders = int(stat.Count)
		}
	}

	// Get product statistics
	var activeProducts int64
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND is_active = ? AND is_approved = ?", sellerID, true, true).
		Count(&activeProducts)
	stats.ActiveProducts = int(activeProducts)

	var inactiveProducts int64
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND (is_active = ? OR is_approved = ?)", sellerID, false, false).
		Count(&inactiveProducts)
	stats.InactiveProducts = int(inactiveProducts)

	var lowStockCount int64
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND quantity < ? AND quantity > ?", sellerID, 10, 0).
		Count(&lowStockCount)
	stats.LowStockProducts = int(lowStockCount)

	var outOfStockCount int64
	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND quantity = ?", sellerID, 0).
		Count(&outOfStockCount)
	stats.OutOfStock = int(outOfStockCount)

	// Get today's statistics
	today := time.Now().Truncate(24 * time.Hour)
	var todayStats struct {
		Sales  float64
		Orders int64
	}
	r.db.Table("order_items").
		Select("COALESCE(SUM(seller_earning), 0) as sales, COUNT(DISTINCT order_id) as orders").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND DATE(orders.created_at) = ?", sellerID, today).
		Scan(&todayStats)
	stats.TodaySales = todayStats.Sales
	stats.TodayOrders = int(todayStats.Orders)

	// Get week's statistics
	weekAgo := time.Now().AddDate(0, 0, -7)
	var weekStats struct {
		Sales  float64
		Orders int64
	}
	r.db.Table("order_items").
		Select("COALESCE(SUM(seller_earning), 0) as sales, COUNT(DISTINCT order_id) as orders").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.created_at >= ?", sellerID, weekAgo).
		Scan(&weekStats)
	stats.WeekSales = weekStats.Sales
	stats.WeekOrders = int(weekStats.Orders)

	// Get month's statistics
	monthAgo := time.Now().AddDate(0, -1, 0)
	var monthStats struct {
		Sales  float64
		Orders int64
	}
	r.db.Table("order_items").
		Select("COALESCE(SUM(seller_earning), 0) as sales, COUNT(DISTINCT order_id) as orders").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.created_at >= ?", sellerID, monthAgo).
		Scan(&monthStats)
	stats.MonthSales = monthStats.Sales
	stats.MonthOrders = int(monthStats.Orders)

	return &stats, nil
}
