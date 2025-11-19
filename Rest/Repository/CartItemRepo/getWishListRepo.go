package cartitemrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *cartRepository) GetWishlist(buyerID uint) ([]models.Wishlist, error) {
	var items []models.Wishlist
	err := r.db.Preload("Product").
		Preload("Product.Seller").
		Preload("Product.Images").
		Where("buyer_id = ?", buyerID).
		Order("created_at DESC").
		Find(&items).Error
	return items, err
}
