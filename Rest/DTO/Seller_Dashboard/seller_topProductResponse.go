package sellerdashboard

// TopProductResponse - Top selling products
type TopProductResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	SKU        string  `json:"sku"`
	Price      float64 `json:"price"`
	TotalSold  int     `json:"total_sold"`
	Revenue    float64 `json:"revenue"`
	ImageURL   string  `json:"image_url"`
	Stock      int     `json:"stock"`
	IsActive   bool    `json:"is_active"`
	IsApproved bool    `json:"is_approved"`
}
