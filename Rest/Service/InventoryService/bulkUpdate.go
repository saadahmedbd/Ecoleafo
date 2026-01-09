package inventoryservice

import (
	"errors"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

// BulkUpdateStock - Update multiple products
func (s *InventoryService) BulkUpdateStock(sellerID uint, updates []inventory.StockUpdate, userID uint) (map[string]interface{}, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	if len(updates) == 0 {
		return nil, errors.New("no updates provided")
	}

	if len(updates) > 100 {
		return nil, errors.New("maximum 100 products can be updated at once")
	}

	successCount := 0
	failedCount := 0
	var errors []string

	for _, update := range updates {
		_, err := s.inventoryRepo.UpdateStock(sellerID, update.ProductID, update.Quantity, "bulk_update", "Bulk stock update", userID)
		if err != nil {
			failedCount++
			errors = append(errors, err.Error())
		} else {
			successCount++
		}
	}

	return map[string]interface{}{
		"success_count": successCount,
		"failed_count":  failedCount,
		"errors":        errors,
	}, nil
}
