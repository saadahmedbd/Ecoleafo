package commissionearningpayoutdto

type SellerEarningDetailsResponse struct {
	SellerID         uint                  `json:"seller_id"`
	SellerName       string                `json:"seller_name"`
	StoreName        string                `json:"store_name"`
	Email            string                `json:"email"`
	Phone            string                `json:"phone"`
	CommissionRate   float64               `json:"commission_rate"`
	TotalOrders      int                   `json:"total_orders"`
	CompletedOrders  int                   `json:"completed_orders"`
	CancelledOrders  int                   `json:"cancelled_orders"`
	GrossSales       float64               `json:"gross_sales"`
	TotalCommission  float64               `json:"total_commission"`
	NetEarnings      float64               `json:"net_earnings"`
	TotalWithdrawn   float64               `json:"total_withdrawn"`
	AvailableBalance float64               `json:"available_balance"`
	PendingClearance float64               `json:"pending_clearance"`
	LastUpdated      string                `json:"last_updated"`
	RecentOrders     []RecentOrderSummary  `json:"recent_orders"`
}

type RecentOrderSummary struct {
	OrderID      uint    `json:"order_id"`
	OrderNumber  string  `json:"order_number"`
	OrderDate    string  `json:"order_date"`
	OrderTotal   float64 `json:"order_total"`
	Commission   float64 `json:"commission"`
	SellerEarned float64 `json:"seller_earned"`
	Status       string  `json:"status"`
}
