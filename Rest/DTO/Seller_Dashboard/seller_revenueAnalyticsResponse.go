package sellerdashboard

// RevenueAnalyticsResponse - Revenue breakdown
type RevenueAnalyticsResponse struct {
	Period       string             `json:"period"`
	TotalRevenue float64            `json:"total_revenue"`
	NetRevenue   float64            `json:"net_revenue"`
	Commission   float64            `json:"commission"`
	Data         []RevenueDataPoint `json:"data"`
	Change       float64            `json:"change"`
}
