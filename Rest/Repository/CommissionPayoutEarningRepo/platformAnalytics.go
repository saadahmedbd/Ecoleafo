package commissionpayoutearningrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

// ============================================================================
// PLATFORM ANALYTICS
// ============================================================================
func (r *commissionRepository) GetPlatformEarningsOverview() (map[string]interface{}, error) {
	var result struct {
		TotalGrossSales       float64
		TotalCommissionEarned float64
		TotalSellerEarnings   float64
		TotalOrders           int64
		CompletedOrders       int64
	}

	// Get totals from order_items table (real-time data)
	r.db.Table("order_items").
		Select("COALESCE(SUM(total), 0) as total_gross_sales, " +
			"COALESCE(SUM(commission), 0) as total_commission_earned, " +
			"COALESCE(SUM(seller_earning), 0) as total_seller_earnings, " +
			"COUNT(DISTINCT order_id) as total_orders").
		Scan(&result)

	// Get completed orders
	r.db.Model(&models.Order{}).
		Where("status = ?", "delivered").
		Count(&result.CompletedOrders)

	// Get payout stats
	var pendingPayouts, completedPayouts float64
	r.db.Model(&models.SellerPayout{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("status = ?", "pending").
		Scan(&pendingPayouts)

	r.db.Model(&models.SellerPayout{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("status = ?", "completed").
		Scan(&completedPayouts)

	return map[string]interface{}{
		"total_gross_sales":       result.TotalGrossSales,
		"total_commission_earned": result.TotalCommissionEarned,
		"total_seller_earnings":   result.TotalSellerEarnings,
		"pending_payouts":         pendingPayouts,
		"completed_payouts":       completedPayouts,
		"total_orders":            result.TotalOrders,
		"completed_orders":        result.CompletedOrders,
	}, nil
}

func (r *commissionRepository) GetMonthlyRevenue(year int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	rows, err := r.db.Table("order_items").
		Select("TO_CHAR(DATE_TRUNC('month', order_items.created_at), 'YYYY-MM') as month, "+
			"COALESCE(SUM(order_items.total), 0) as gross_revenue, "+
			"COALESCE(SUM(order_items.commission), 0) as commission_earned, "+
			"COUNT(DISTINCT order_items.order_id) as total_orders").
		Where("EXTRACT(YEAR FROM order_items.created_at) = ?", year).
		Group("DATE_TRUNC('month', order_items.created_at)").
		Order("DATE_TRUNC('month', order_items.created_at) ASC").
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var month string
		var grossRevenue, commissionEarned float64
		var totalOrders int

		if err := rows.Scan(&month, &grossRevenue, &commissionEarned, &totalOrders); err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"month":             month,
			"gross_revenue":     grossRevenue,
			"commission_earned": commissionEarned,
			"total_orders":      totalOrders,
		})
	}

	return results, nil
}

func (r *commissionRepository) GetTopSellersByRevenue(limit int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	rows, err := r.db.Table("order_items").
		Select("order_items.seller_id, u.store_name as seller_name, "+
			"COUNT(DISTINCT order_items.order_id) as total_orders, "+
			"COALESCE(SUM(order_items.total), 0) as gross_sales, "+
			"COALESCE(SUM(order_items.commission), 0) as commission_generated").
		Joins("JOIN users u ON u.id = order_items.seller_id").
		Group("order_items.seller_id, u.store_name").
		Order("gross_sales DESC").
		Limit(limit).
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sellerID uint
		var sellerName string
		var totalOrders int
		var grossSales, commissionGenerated float64

		if err := rows.Scan(&sellerID, &sellerName, &totalOrders, &grossSales, &commissionGenerated); err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"seller_id":            sellerID,
			"seller_name":          sellerName,
			"total_orders":         totalOrders,
			"gross_sales":          grossSales,
			"commission_generated": commissionGenerated,
		})
	}

	return results, nil
}
