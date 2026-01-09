package sellerdashboard

// SalesAnalyticsResponse - Sales data for charts
type SalesAnalyticsResponse struct {
	Period string           `json:"period"` // week, month, year
	Data   []SalesDataPoint `json:"data"`
	Total  float64          `json:"total"`
	Change float64          `json:"change"` // percentage change from previous period
}
