package inventoryhandler

import inventoryservice "github.com/saadahmedbd/Treestore/Rest/Service/InventoryService"

type Inventoryhandler struct {
	inventoryservice *inventoryservice.InventoryService
}

func NewInventoryHandler(inventoryservice *inventoryservice.InventoryService) *Inventoryhandler {
	return &Inventoryhandler{
		inventoryservice: inventoryservice,
	}
}
