package sellerdashboard

// OrderStatusDistribution - Order status breakdown
type OrderStatusDistribution struct {
	Status string  `json:"status"`
	Count  int     `json:"count"`
	Color  string  `json:"color"`
	Value  float64 `json:"value"` // for pie chart
}
