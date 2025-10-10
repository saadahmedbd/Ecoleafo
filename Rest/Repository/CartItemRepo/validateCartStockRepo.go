package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) ValidateCartStock(buyerID uint) ([]uint, error) {
	var invalidProductIDs []uint

	var cartItems []models.CartItem
	r.db.Preload("Product").Where("buyer_id = ? AND is_saved_for_later = ?", buyerID, false).
		Find(&cartItems)

	for _, item := range cartItems {
		if !item.Product.IsActive || item.Product.Quantity < item.Quantity {
			invalidProductIDs = append(invalidProductIDs, item.ProductID)
		}
	}

	return invalidProductIDs, nil
}
