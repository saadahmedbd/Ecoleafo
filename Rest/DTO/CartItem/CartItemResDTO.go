package cartitem

type CartItemResponse struct {
	ID              uint     `json:"id"`
	ProductID       uint     `json:"product_id"`
	ProductName     string   `json:"product_name"`
	ProductSlug     string   `json:"product_slug"`
	Price           float64  `json:"price"`
	DiscountPrice   float32  `json:"discount_price"`
	OriginalPrice   float64  `json:"original_price"`
	DiscountPercent float64  `json:"discount_percent"`
	Quantity        int      `json:"quantity"`
	Subtotal        float64  `json:"subtotal"`
	Image           []string `json:"image"`
	SellerName      string   `json:"seller_name"`
	SellerID        uint     `json:"seller_id"`

	// Stock & Availability
	InStock         bool   `json:"in_stock"`
	StockQuantity   int    `json:"stock_quantity"`
	IsAvailable     bool   `json:"is_available"`
	AvailabilityMsg string `json:"availability_message"`

	// Additional features
	IsGift          bool   `json:"is_gift"`
	GiftMessage     string `json:"gift_message"`
	IsSavedForLater bool   `json:"is_saved_for_later"`
	IsSelected      bool   `json:"is_selected"`

	// Recommendations
	CanIncreaseQty bool `json:"can_increase_qty"`
	MaxQuantity    int  `json:"max_quantity"`
}
