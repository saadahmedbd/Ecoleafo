package inventoryrepo

import (
	"errors"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

func (r *InventoryRepository) UpdateStock(sellerID uint, productID uint, newQty int, reason, ref string, userID uint) (*inventory.InventoryItemResponse, error) {
	var product models.Product
	if err := r.db.Model(&models.Product{}).Where("id = ?", productID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	if product.SellerID != sellerID {
		return nil, errors.New("unauthorized: product belongs to different seller")
	}

	prevStock := product.Quantity

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&product).Update("quantity", newQty).Error; err != nil {
			return err
		}

		history := models.StockHistory{
			ProductID: productID,
			SellerID:  sellerID,
			Type:      "adjustment",
			Quantity:  newQty - prevStock,
			PrevStock: prevStock,
			NewStock:  newQty,
			Reason:    reason,
			Reference: ref,
			CreatedBy: userID,
			CreatedAt: time.Now(),
		}
		return tx.Create(&history).Error
	})

	if err != nil {
		return nil, err
	}

	return r.GetInventoryItem(productID)
}

// func (r *InventoryRepository) UpdateStock(
// 	sellerID uint,
// 	productID uint,
// 	newQuantity int,
// 	changedBy uint, // user/admin ID who made the change
// 	reason string, // optional reason
// 	reference string, // e.g., order_id, adjustment_id
// ) (*inventory.InventoryItemResponse, error) {

// 	var item inventory.InventoryItemResponse

// 	err := r.db.Transaction(func(tx *gorm.DB) error {
// 		// 1. Update stock and fetch updated product with primary image
// 		err := tx.Raw(`
// 			SELECT
// 				p.id, p.name, p.sku, p.category_id,
// 				c.name AS category_name, p.price, p.quantity AS stock,
// 				p.min_quantity AS low_stock_threshold, p.is_active, p.is_approved,
// 				p.updated_at AS last_updated,
// 				(p.quantity * p.price) AS stock_value,
// 				pi.image_url AS image_url,
// 				p.quantity AS prev_quantity
// 			FROM (
// 				UPDATE products
// 				SET quantity = ?
// 				WHERE id = ? AND seller_id = ?
// 				RETURNING *
// 			) p
// 			LEFT JOIN categories c ON p.category_id = c.id
// 			LEFT JOIN LATERAL (
// 				SELECT image_url
// 				FROM product_images
// 				WHERE product_id = p.id
// 				ORDER BY is_primary DESC, sort_order ASC
// 				LIMIT 1
// 			) pi ON true
// 		`, newQuantity, productID, sellerID).Scan(&item).Error
// 		if err != nil {
// 			return err
// 		}

// 		if item.ID == 0 {
// 			return errors.New("product not found or unauthorized")
// 		}

// 		prevStock := item.Stock
// 		if prevStock == newQuantity {
// 			// No change; skip logging
// 			return nil
// 		}

// 		// 2. Determine change type
// 		changeType := "adjustment"
// 		if newQuantity < prevStock {
// 			changeType = "sale"
// 		} else if newQuantity > prevStock {
// 			changeType = "restock"
// 		}

// 		// 3. Log stock change
// 		history := models.StockHistory{
// 			ProductID: productID,
// 			SellerID:  sellerID,
// 			Type:      changeType,
// 			Quantity:  newQuantity - prevStock,
// 			PrevStock: prevStock,
// 			NewStock:  newQuantity,
// 			Reason:    reason,
// 			Reference: reference,
// 			CreatedBy: changedBy,
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		}

// 		return tx.Create(&history).Error
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &item, nil
// }
