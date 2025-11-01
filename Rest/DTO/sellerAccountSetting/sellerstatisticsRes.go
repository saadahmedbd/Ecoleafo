package selleraccountsetting

// SellerStatisticsResponse - Seller statistics
type SellerStatisticsResponse struct {
	TotalSales      float64 `json:"total_sales"`
	TotalOrders     int     `json:"total_orders"`
	ActiveProducts  int     `json:"active_products"`
	Rating          float64 `json:"rating"`
	TotalReviews    int     `json:"total_reviews"`
	PendingOrders   int     `json:"pending_orders"`
	CompletedOrders int     `json:"completed_orders"`
	TotalEarnings   float64 `json:"total_earnings"`
}
