package sellerdashboard

// PerformanceMetrics - Performance indicators
type PerformanceMetrics struct {
	SuccessRate          float64 `json:"success_rate"`
	CustomerSatisfaction float64 `json:"customer_satisfaction"`
	AverageOrderValue    float64 `json:"average_order_value"`
	FulfillmentRate      float64 `json:"fulfillment_rate"`
	ReturnRate           float64 `json:"return_rate"`
	ConversionRate       float64 `json:"conversion_rate"`
}
