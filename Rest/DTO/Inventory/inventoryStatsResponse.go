package inventory

// InventoryStatsResponse - Inventory statistics
type InventoryStatsResponse struct {
	TotalStock      int     `json:"total_stock"`        // Total quantity across all products
	LowStockCount   int     `json:"low_stock_count"`    // Products below threshold
	OutOfStockCount int     `json:"out_of_stock_count"` // Products with 0 stock
	TotalValue      float64 `json:"total_value"`        // Total inventory value
	TotalProducts   int     `json:"total_products"`     // Total number of products
	AverageStock    float64 `json:"average_stock"`      // Average stock per product
}
