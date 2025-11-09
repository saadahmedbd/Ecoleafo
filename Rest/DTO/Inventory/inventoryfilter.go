package inventory

// InventoryFilter - Query parameters
type InventoryFilter struct {
	Search   string `json:"search"`
	Status   string `json:"status"` // all, low-stock, out-of-stock, in-stock
	Category string `json:"category"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}
