package cartitem

type CartSummaryResponse struct {
	Items         []CartItemResponse `json:"items"`
	SavedForLater []CartItemResponse `json:"saved_for_later"`

	// Pricing
	ItemCount      int     `json:"item_count"`
	SavedItemCount int     `json:"saved_item_count"`
	Subtotal       float64 `json:"subtotal"`
	Discount       float64 `json:"discount"`
	Tax            float64 `json:"tax"`
	ShippingCost   float64 `json:"shipping_cost"`
	TotalAmount    float64 `json:"total_amount"`
	TotalSavings   float64 `json:"total_savings"`

	// Coupon
	CouponApplied  bool    `json:"coupon_applied"`
	CouponCode     string  `json:"coupon_code"`
	CouponDiscount float64 `json:"coupon_discount"`

	// Availability
	HasUnavailableItems bool     `json:"has_unavailable_items"`
	UnavailableItems    []string `json:"unavailable_items"`
	CanCheckout         bool     `json:"can_checkout"`

	// Recommendations
	RecommendedProducts []ProductRecommendation `json:"recommended_products"`

	// Messages
	Messages []CartMessage `json:"messages"`
}
