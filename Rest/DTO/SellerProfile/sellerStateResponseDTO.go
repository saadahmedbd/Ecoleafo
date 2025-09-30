package sellerprofile

type SellerStateResponse struct {
	TotalSales      float64 `json:"total_sales"`
	TotalEarnings   float64 `json:"total_earnings"`
	TotalOrders     int     `json:"total_orders"`
	TotalProducts   int     `json:"total_products"`
	ActiveProducts  int     `json:"active_products"`
	PendingProducts int     `json:"pending_products"`
	AverageRating   float64 `json:"average_rating"`
	ReviewCount     int     `json:"review_count"`
}
