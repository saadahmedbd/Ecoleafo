package inventory

import "time"

// InventoryItemResponse - Single inventory item
type InventoryItemResponse struct {
	ID                uint      `json:"id"`
	Name              string    `json:"name"`
	SKU               string    `json:"sku"`
	CategoryID        uint      `json:"category_id"`
	CategoryName      string    `json:"category_name"`
	Price             float64   `json:"price"`
	Stock             int       `json:"stock"`
	LowStockThreshold int       `json:"low_stock_threshold"`
	ImageURL          string    `json:"image_url"`
	IsActive          bool      `json:"is_active"`
	IsApproved        bool      `json:"is_approved"`
	LastUpdated       time.Time `json:"last_updated"`
	StockValue        float64   `json:"stock_value"` // stock * price
	TotalCount        uint      `json:"total_count"`
}
