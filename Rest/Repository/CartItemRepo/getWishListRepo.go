package cartitemrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) GetWishlist(buyerID uint) ([]models.Wishlist, error) {
	var items []models.Wishlist
	err := r.db.Preload("Product").
		Preload("Product.Images", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_primary = ?", true).Limit(1)
		}).
		Where("buyer_id = ?", buyerID).
		Order("created_at DESC").
		Find(&items).Error
	return items, err
}
