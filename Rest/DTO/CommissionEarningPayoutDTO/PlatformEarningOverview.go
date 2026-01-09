package commissionearningpayoutdto

type PlatformEarningsOverview struct {
	TotalGrossSales       float64 `json:"total_gross_sales"`
	TotalCommissionEarned float64 `json:"total_commission_earned"`
	TotalSellerEarnings   float64 `json:"total_seller_earnings"`
	PendingPayouts        float64 `json:"pending_payouts"`
	CompletedPayouts      float64 `json:"completed_payouts"`
	TotalOrders           int     `json:"total_orders"`
	CompletedOrders       int     `json:"completed_orders"`
}
