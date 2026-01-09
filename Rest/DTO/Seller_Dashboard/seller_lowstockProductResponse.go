package sellerdashboard

// LowStockProductResponse - Products with low stock
type LowStockProductResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Quantity int     `json:"quantity"`
	MinStock int     `json:"min_stock"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
	Status   string  `json:"status"` // low-stock, out-of-stock
	IsActive bool    `json:"is_active"`
}
