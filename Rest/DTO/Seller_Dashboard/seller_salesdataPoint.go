package sellerdashboard

import "time"

// SalesDataPoint - Single data point for charts
type SalesDataPoint struct {
	Name    string    `json:"name"` // Mon, Tue, Jan, Feb, etc.
	Date    time.Time `json:"date"`
	Sales   float64   `json:"sales"`
	Orders  int       `json:"orders"`
	Revenue float64   `json:"revenue"`
}
