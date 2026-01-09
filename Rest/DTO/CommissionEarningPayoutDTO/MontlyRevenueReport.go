package commissionearningpayoutdto

type MonthlyRevenueReport struct {
	Month            string  `json:"month"`
	GrossRevenue     float64 `json:"gross_revenue"`
	CommissionEarned float64 `json:"commission_earned"`
	TotalOrders      int     `json:"total_orders"`
}
