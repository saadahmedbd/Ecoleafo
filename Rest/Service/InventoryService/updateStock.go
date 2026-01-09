package inventoryservice

import (
	"errors"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

// UpdateStock - Update product stock
func (s *InventoryService) UpdateStock(sellerID uint, productID uint, quantity int, userID uint) (*inventory.InventoryItemResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	if productID == 0 {
		return nil, errors.New("invalid product ID")
	}

	if quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}

	// Update stock
	product, err := s.inventoryRepo.UpdateStock(sellerID, productID, quantity, "update", " stock update", userID)
	if err != nil {
		return nil, err
	}

	// Create stock history entry
	err = s.inventoryRepo.CreateStockHistory(productID, quantity, "manual", "Stock updated manually")
	if err != nil {
		// Log error but don't fail the operation
	}

	return product, nil
}
