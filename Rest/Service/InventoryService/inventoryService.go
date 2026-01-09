package inventoryservice

import (
	inventoryrepo "github.com/saadahmedbd/Treestore/Rest/Repository/InventoryRepo"
	sellerrepo "github.com/saadahmedbd/Treestore/Rest/Repository/sellerAccountSettingRepo"
)

type InventoryService struct {
	inventoryRepo *inventoryrepo.InventoryRepository
	sellerRepo    *sellerrepo.SellerAccountSettingRepository
}

func NewInventoryService(inventoryRepo *inventoryrepo.InventoryRepository, sellerRepo *sellerrepo.SellerAccountSettingRepository) *InventoryService {
	return &InventoryService{
		inventoryRepo: inventoryRepo,
		sellerRepo:    sellerRepo,
	}
}
