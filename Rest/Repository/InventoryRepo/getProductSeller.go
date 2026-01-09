package inventoryrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetProductBySeller - Verify product ownership
func (r *InventoryRepository) GetProductBySeller(sellerID uint, productID uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND seller_id = ?", productID, sellerID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}
