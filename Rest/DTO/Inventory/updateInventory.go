package inventory

// UpdateStockRequest - Update single product stock
type UpdateStockRequest struct {
	Quantity int `json:"quantity" validate:"required,gte=0"`
}

// UpdateThresholdRequest - Update low stock threshold
type UpdateThresholdRequest struct {
	MinQuantity int `json:"min_quantity" validate:"required,gte=0"`
}

// BulkUpdateStockRequest - Bulk stock update
type BulkUpdateStockRequest struct {
	Updates []StockUpdate `json:"updates" validate:"required,dive"`
}

// StockUpdate - Single stock update entry
type StockUpdate struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gte=0"`
}
