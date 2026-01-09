package sellerdashboardrepo

import (
	"time"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetSalesAnalytics - Get sales data for charts
func (r *DashboardRepository) GetSalesAnalytics(regUserID uint, period string) (*sellerdashboard.SalesAnalyticsResponse, error) {
	// Convert reguser ID to seller ID
	var seller struct {
		ID uint
	}
	if err := r.db.Table("users").Select("id").Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return nil, err
	}
	sellerID := seller.ID

	var dataPoints []sellerdashboard.SalesDataPoint
	var startDate time.Time
	var groupBy string

	now := time.Now()

	switch period {
	case "week":
		startDate = now.AddDate(0, 0, -7)
		groupBy = "DATE(orders.created_at)"
	case "month":
		startDate = now.AddDate(0, -1, 0)
		groupBy = "DATE(orders.created_at)"
	case "year":
		startDate = now.AddDate(-1, 0, 0)
		groupBy = "DATE_TRUNC('month', orders.created_at)"
	default:
		startDate = now.AddDate(0, 0, -7)
		groupBy = "DATE(orders.created_at)"
	}

	// Query sales data grouped by period
	type SalesGroup struct {
		Date   time.Time
		Sales  float64
		Orders int64
	}
	var salesGroups []SalesGroup

	err := r.db.Table("order_items").
		Select(groupBy+" as date, COALESCE(SUM(seller_earning), 0) as sales, COUNT(DISTINCT order_id) as orders").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.created_at >= ?", sellerID, startDate).
		Group(groupBy).
		Order("date ASC").
		Scan(&salesGroups).Error

	if err != nil {
		return nil, err
	}

	// Format data points
	var total float64
	for _, group := range salesGroups {
		name := group.Date.Format("Mon")
		if period == "month" {
			name = group.Date.Format("Jan 02")
		} else if period == "year" {
			name = group.Date.Format("Jan")
		}

		dataPoints = append(dataPoints, sellerdashboard.SalesDataPoint{
			Name:    name,
			Date:    group.Date,
			Sales:   group.Sales,
			Orders:  int(group.Orders),
			Revenue: group.Sales,
		})
		total += group.Sales
	}

	// Calculate change from previous period
	var previousTotal float64
	previousStart := startDate.AddDate(0, 0, -int(now.Sub(startDate).Hours()/24))
	r.db.Table("order_items").
		Select("COALESCE(SUM(seller_earning), 0)").
		Joins("JOIN orders ON order_items.order_id = orders.id").
		Where("order_items.seller_id = ? AND orders.created_at >= ? AND orders.created_at < ?", sellerID, previousStart, startDate).
		Scan(&previousTotal)

	change := 0.0
	if previousTotal > 0 {
		change = ((total - previousTotal) / previousTotal) * 100
	}

	return &sellerdashboard.SalesAnalyticsResponse{
		Period: period,
		Data:   dataPoints,
		Total:  total,
		Change: change,
	}, nil
}
