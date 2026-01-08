package commissionearningpayoutdto

type SellerEarningsResponse struct {
	SellerID         uint    `json:"seller_id"`
	SellerName       string  `json:"seller_name"`
	StoreName        string  `json:"store_name"`
	CommissionRate   float64 `json:"commission_rate"`
	TotalOrders      int     `json:"total_orders"`
	CompletedOrders  int     `json:"completed_orders"`
	GrossSales       float64 `json:"gross_sales"`
	TotalCommission  float64 `json:"total_commission"`
	NetEarnings      float64 `json:"net_earnings"`
	TotalWithdrawn   float64 `json:"total_withdrawn"`
	AvailableBalance float64 `json:"available_balance"`
	PendingClearance float64 `json:"pending_clearance"`
	LastUpdated      string  `json:"last_updated"`
}
