package inventoryservice

import (
	"errors"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

// GetInventoryStats - Get inventory statistics
func (s *InventoryService) GetInventoryStats(sellerID uint) (*inventory.InventoryStatsResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	stats, err := s.inventoryRepo.GetInventoryStats(sellerID)
	if err != nil {
		return nil, errors.New("failed to fetch statistics")
	}

	return stats, nil
}
