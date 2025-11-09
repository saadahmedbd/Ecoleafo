package inventoryservice

import (
	"errors"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

// GetStockHistory - Get stock change history
func (s *InventoryService) GetStockHistory(sellerID uint, productID uint, limit int) ([]inventory.StockHistoryResponse, error) {
	if sellerID == 0 {
		return nil, errors.New("invalid seller ID")
	}

	if productID == 0 {
		return nil, errors.New("invalid product ID")
	}

	// Verify seller owns product
	_, err := s.inventoryRepo.GetProductBySeller(sellerID, productID)
	if err != nil {
		return nil, errors.New("product not found or unauthorized")
	}

	history, err := s.inventoryRepo.GetStockHistory(productID, sellerID, limit)
	if err != nil {
		return nil, errors.New("failed to fetch history")
	}

	return history, nil
}
