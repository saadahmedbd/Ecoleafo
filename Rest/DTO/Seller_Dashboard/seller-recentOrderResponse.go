package sellerdashboard

import "time"

// RecentOrderResponse - Recent orders for dashboard
type RecentOrderResponse struct {
	ID           uint      `json:"id"`
	OrderNumber  string    `json:"order_number"`
	CustomerName string    `json:"customer_name"`
	ProductName  string    `json:"product_name"`
	Total        float64   `json:"total"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	TimeAgo      string    `json:"time_ago"`
}
