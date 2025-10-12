package order

type OrderListFilter struct {
	Status        string `json:"status"`
	PaymentStatus string `json:"payment_status"`
	BuyerID       uint   `json:"buyer_id"`
	SellerID      uint   `json:"seller_id"`
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
}
