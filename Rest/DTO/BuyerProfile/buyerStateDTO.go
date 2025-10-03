package buyerprofile

import "time"

type BuyerStateResponse struct {
	TotalOrders     int               `json:"total_orders"`
	PendingOrders   int               `json:"pending_orders"`
	CompletedOrders int               `json:"completed_orders"`
	CancelledOrders int               `json:"cancelled_orders"`
	TotalSpent      float64           `json:"total_spent"`
	WishlistCount   int               `json:"wishlist_count"`
	CartItemCount   int               `json:"cart_item_count"`
	ReviewCount     int               `json:"review_count"`
	LastOrderDate   *time.Time        `json:"last_order_date"`
	RecentOrders    []RecentOrderInfo `json:"recent_orders"`
}
