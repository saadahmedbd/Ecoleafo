package inventoryrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

func (r *InventoryRepository) GetInventoryStats(sellerID uint) (*inventory.InventoryStatsResponse, error) {
	var result struct {
		TotalProducts   int64
		TotalStock      int64
		TotalValue      float64
		LowStockCount   int64
		OutOfStockCount int64
	}

	// Single query with conditional SUMs
	err := r.db.Model(&models.Product{}).
		Select(`COUNT(*) as total_products,
				COALESCE(SUM(quantity),0) as total_stock,
				COALESCE(SUM(quantity * price),0) as total_value,
				COALESCE(SUM(CASE WHEN quantity > 0 AND quantity <= min_quantity THEN 1 ELSE 0 END),0) as low_stock_count,
				COALESCE(SUM(CASE WHEN quantity = 0 THEN 1 ELSE 0 END),0) as out_of_stock_count`).
		Where("seller_id = ?", sellerID).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	stats := &inventory.InventoryStatsResponse{
		TotalProducts:   int(result.TotalProducts),
		TotalStock:      int(result.TotalStock),
		TotalValue:      result.TotalValue,
		LowStockCount:   int(result.LowStockCount),
		OutOfStockCount: int(result.OutOfStockCount),
	}

	// Compute average stock safely
	if stats.TotalProducts > 0 {
		stats.AverageStock = float64(stats.TotalStock) / float64(stats.TotalProducts)
	}

	return stats, nil
}
