package inventory

// InventoryListResponse - List of inventory items
type InventoryListResponse struct {
	Items []InventoryItemResponse `json:"items"`
	Total int64                   `json:"total"`
	Page  int                     `json:"page"`
	Limit int                     `json:"limit"`
}
