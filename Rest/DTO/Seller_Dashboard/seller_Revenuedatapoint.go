package sellerdashboard

import "time"

// RevenueDataPoint - Single revenue data point
type RevenueDataPoint struct {
	Name       string    `json:"name"`
	Date       time.Time `json:"date"`
	Revenue    float64   `json:"revenue"`
	Commission float64   `json:"commission"`
	Net        float64   `json:"net"`
}
