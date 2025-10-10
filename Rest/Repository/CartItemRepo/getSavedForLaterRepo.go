package cartitemrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) GetSavedForLater(buyerID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.db.Preload("Product").
		Preload("Product.Images", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_primary = ?", true).Limit(1)
		}).
		Where("buyer_id = ? AND is_saved_for_later = ?", buyerID, true).
		Find(&items).Error
	return items, err
}
