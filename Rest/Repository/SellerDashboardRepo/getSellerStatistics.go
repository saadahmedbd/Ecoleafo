package sellerdashboardrepo

import (
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetSellerStatistics - Get comprehensive seller statistics
func (r *DashboardRepository) GetSellerStatistics(regUserID uint) (*sellerdashboard.DashboardStatsResponse, error) {
	var stats sellerdashboard.DashboardStatsResponse

	// Convert reguser ID to seller ID
	var seller models.User
	if err := r.db.Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return nil, fmt.Errorf("seller not found for reguser_id %d", regUserID)
	}

	sellerID := seller.ID
	stats.AverageRating = seller.AverageRating
	stats.TotalReviews = seller.TotalReviews

	// Calculate total sales and earnings from order_items (real-time)
	var totals struct {
		TotalSales      float64
		TotalEarnings   float64
		TotalCommission float64
		TotalOrders     int64
	}
	r.db.Table("order_items").
		Select("COALESCE(SUM(total), 0) as total_sales, COALESCE(SUM(seller_earning), 0) as total_earnings, COALESCE(SUM(commission), 0) as total_commission, COUNT(DISTINCT order_id) as total_orders").
		Where("seller_id = ?", sellerID).
		Scan(&totals)

	stats.TotalSales = totals.TotalSales
	stats.TotalEarnings = totals.TotalEarnings
	stats.TotalCommission = totals.TotalCommission
	stats.TotalOrders = int(totals.TotalOrders)

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
	result := r.db.Model(&models.Product{}).
		Where("seller_id = ? AND is_active = ? AND is_approved = ?", sellerID, true, true).
		Count(&activeProducts)
	fmt.Printf("[DEBUG] Active products query - SellerID: %d, Count: %d, Error: %v\n", sellerID, activeProducts, result.Error)
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
	result2 := r.db.Table("order_items").
		Select("COALESCE(SUM(seller_earning), 0) as sales, COUNT(DISTINCT order_id) as orders").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND DATE(orders.created_at) = ?", sellerID, today).
		Scan(&todayStats)
	fmt.Printf("[DEBUG] Today stats - SellerID: %d, Sales: %.2f, Orders: %d, Error: %v\n", sellerID, todayStats.Sales, todayStats.Orders, result2.Error)
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
