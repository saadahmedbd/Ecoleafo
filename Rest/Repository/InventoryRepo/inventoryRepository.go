package inventoryrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (r *InventoryRepository) LogStockChange(
	sellerID uint,
	productID uint,
	changeType string, // "restock", "sale", "adjustment", "return"
	quantity int,
	prevStock int,
	newStock int,
	reason string,
	reference string,
	createdBy uint,
) error {
	history := models.StockHistory{
		ProductID: productID,
		SellerID:  sellerID,
		Type:      changeType,
		Quantity:  quantity,
		PrevStock: prevStock,
		NewStock:  newStock,
		Reason:    reason,
		Reference: reference,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return r.db.Create(&history).Error
}
