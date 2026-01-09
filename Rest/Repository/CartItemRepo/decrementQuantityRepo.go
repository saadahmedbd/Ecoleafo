package cartitemrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *cartRepository) DecrementQuantity(buyerID, productID uint) error {
	return r.db.Model(&models.CartItem{}).
		Where("buyer_id = ? AND product_id = ? AND quantity > ?", buyerID, productID, 1).
		UpdateColumn("quantity", gorm.Expr("quantity - ?", 1)).Error
}
