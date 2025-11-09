package inventoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CreateStockHistory - Create stock history entry
func (r *InventoryRepository) CreateStockHistory(productID uint, newQuantity int, changeType string, reason string) error {
	// Get current quantity
	var product models.Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return err
	}

	// Create history entry (you can create a StockHistory model if needed)
	// For now, we're just tracking in product updates
	// You can extend this by creating a stock_history table

	return nil
}
