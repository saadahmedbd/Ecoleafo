package cartitemrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) GetCart(buyerID uint, includeSavedForLater bool) ([]models.CartItem, error) {
	var items []models.CartItem
	query := r.db.Preload("Product").
		Preload("Product.Images", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_primary = ?", true).Limit(1)
		}).
		Preload("Product.Seller", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, store_name")
		}).
		Where("buyer_id = ?", buyerID)

	if !includeSavedForLater {
		query = query.Where("is_saved_for_later = ?", false)
	}

	err := query.Order("created_at DESC").Find(&items).Error
	return items, err
}
