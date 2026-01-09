package sellerdashboard

// DashboardStatsResponse - Main dashboard statistics
type DashboardStatsResponse struct {
	CommissionRate   float64 `json:"commission_rate"`
	TotalSales       float64 `json:"total_sales"`
	TotalOrders      int     `json:"total_orders"`
	TotalEarnings    float64 `json:"total_earnings"`
	TotalCommission  float64 `json:"total_commission"`
	AverageRating    float64 `json:"average_rating"`
	TotalReviews     int     `json:"total_reviews"`
	PendingOrders    int     `json:"pending_orders"`
	CompletedOrders  int     `json:"completed_orders"`
	ShippedOrders    int     `json:"shipped_orders"`
	CancelledOrders  int     `json:"cancelled_orders"`
	ActiveProducts   int     `json:"active_products"`
	InactiveProducts int     `json:"inactive_products"`
	LowStockProducts int     `json:"low_stock_products"`
	OutOfStock       int     `json:"out_of_stock"`
	TodaySales       float64 `json:"today_sales"`
	TodayOrders      int     `json:"today_orders"`
	WeekSales        float64 `json:"week_sales"`
	WeekOrders       int     `json:"week_orders"`
	MonthSales       float64 `json:"month_sales"`
	MonthOrders      int     `json:"month_orders"`
}
