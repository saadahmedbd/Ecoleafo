package guestcartitem

type CartItemInfo struct {
	ID          uint    `json:"id"`
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	ProductSlug string  `json:"product_slug"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
	Image       string  `json:"image"`
	InStock     bool    `json:"in_stock"`
}
