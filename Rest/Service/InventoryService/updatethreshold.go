package inventoryservice

import (
	"errors"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

// UpdateThreshold - Update low stock threshold
func (s *InventoryService) UpdateThreshold(sellerID uint, productID uint, minQuantity int) (*inventory.InventoryItemResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	if productID == 0 {
		return nil, errors.New("invalid product ID")
	}

	if minQuantity < 0 {
		return nil, errors.New("threshold cannot be negative")
	}

	product, err := s.inventoryRepo.UpdateThreshold(sellerID, productID, minQuantity)
	if err != nil {
		return nil, err
	}

	return product, nil
}
