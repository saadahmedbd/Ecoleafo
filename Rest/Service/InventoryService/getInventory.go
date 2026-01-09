package inventoryservice

import (
	"errors"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

// GetInventory - Get inventory with filters
func (s *InventoryService) GetInventory(userID uint, filters inventory.InventoryFilter) (*inventory.InventoryListResponse, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	// Get seller by RegUser ID (from JWT)
	seller, err := s.sellerRepo.GetSellerByUserID(userID)
	if err != nil {
		return nil, errors.New("seller not found")
	}

	// Set defaults
	if filters.Page == 0 {
		filters.Page = 1
	}
	if filters.Limit == 0 {
		filters.Limit = 50
	}

	// Use actual seller ID for inventory query
	inventoryList, err := s.inventoryRepo.GetInventory(seller.ID, filters)
	if err != nil {
		return nil, errors.New("failed to fetch inventory")
	}

	return inventoryList, nil
}
