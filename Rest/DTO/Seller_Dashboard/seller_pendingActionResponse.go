package sellerdashboard

// PendingActionsResponse - Actions requiring attention
type PendingActionsResponse struct {
	PendingOrders   int `json:"pending_orders"`
	PendingProducts int `json:"pending_products"`
	OutOfStock      int `json:"out_of_stock"`
	LowStock        int `json:"low_stock"`
	PendingReviews  int `json:"pending_reviews"`
	UnreadMessages  int `json:"unread_messages"`
}
