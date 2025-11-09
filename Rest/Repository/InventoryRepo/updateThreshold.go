package inventoryrepo

import (
	"errors"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *InventoryRepository) UpdateThreshold(sellerID uint, productID uint, minQty int) (*inventory.InventoryItemResponse, error) {
	// Transaction ensures consistency between threshold + product
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var product models.Product
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ? AND seller_id = ?", productID, sellerID).
		First(&product).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found or unauthorized")
		}
		return nil, err
	}

	if err := tx.Model(&product).Update("min_quantity", minQty).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Record or update threshold table (for analytics)
	var threshold models.InventoryThreshold
	err := tx.Where("seller_id = ? AND product_id = ?", sellerID, productID).First(&threshold).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		threshold = models.InventoryThreshold{
			SellerID:    sellerID,
			ProductID:   productID,
			CategoryID:  product.CategoryID,
			MinQuantity: minQty,
			AlertLevel:  "normal",
			CreatedAt:   time.Now(),
		}
		if err := tx.Create(&threshold).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	} else if err == nil {
		tx.Model(&threshold).Update("min_quantity", minQty)
	}

	tx.Commit()
	return r.GetInventoryItem(productID)
}

//fully optimize zero redundet but not for grwoing team

// func (r *InventoryRepository) UpdateThreshold(
// 	sellerID uint,
// 	productID uint,
// 	minQty int,
// ) (*dto.InventoryItemResponse, error) {

// 	var item dto.InventoryItemResponse

// 	err := r.db.Transaction(func(tx *gorm.DB) error {
// 		// 1️⃣ Update product threshold with lock and return data
// 		err := tx.Raw(`
// 			WITH updated AS (
// 				UPDATE products
// 				SET min_quantity = @minQty, updated_at = NOW()
// 				WHERE id = @productID AND seller_id = @sellerID
// 				RETURNING id, name, sku, category_id, price, quantity, min_quantity, is_active, is_approved, updated_at
// 			)
// 			SELECT
// 				u.id, u.name, u.sku, u.category_id,
// 				c.name AS category_name,
// 				u.price, u.quantity AS stock,
// 				u.min_quantity AS low_stock_threshold,
// 				u.is_active, u.is_approved,
// 				u.updated_at AS last_updated,
// 				(u.quantity * u.price) AS stock_value,
// 				pi.image_url AS image_url
// 			FROM updated u
// 			LEFT JOIN categories c ON u.category_id = c.id
// 			LEFT JOIN LATERAL (
// 				SELECT image_url
// 				FROM product_images
// 				WHERE product_id = u.id
// 				ORDER BY is_primary DESC, sort_order ASC
// 				LIMIT 1
// 			) pi ON TRUE
// 		`, map[string]interface{}{
// 			"minQty":    minQty,
// 			"productID": productID,
// 			"sellerID":  sellerID,
// 		}).Scan(&item).Error

// 		if err != nil {
// 			return err
// 		}

// 		if item.ID == 0 {
// 			return errors.New("product not found or unauthorized")
// 		}

// 		// 2️⃣ Upsert threshold record efficiently
// 		err = tx.Exec(`
// 			INSERT INTO inventory_thresholds (seller_id, product_id, min_quantity, alert_level, created_at, updated_at)
// 			VALUES (?, ?, ?, 'normal', NOW(), NOW())
// 			ON CONFLICT (seller_id, product_id)
// 			DO UPDATE SET min_quantity = EXCLUDED.min_quantity, updated_at = NOW()
// 		`, sellerID, productID, minQty).Error

// 		return err
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &item, nil
// }
