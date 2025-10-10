package cartitem

type ProductRecommendation struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Reason      string  `json:"reason"` // "frequently_bought_together", "similar_items"
}
