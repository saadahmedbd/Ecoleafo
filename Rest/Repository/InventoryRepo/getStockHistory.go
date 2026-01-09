package inventoryrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

func (r *InventoryRepository) GetStockHistory(productID, sellerID uint, limit int) ([]inventory.StockHistoryResponse, error) {
	if limit <= 0 {
		limit = 50
	}

	var history []inventory.StockHistoryResponse
	err := r.db.Model(&models.StockHistory{}).
		Select("id, product_id, seller_id, type, quantity, prev_stock, new_stock, reason, reference, created_by, created_at").
		Where("product_id = ? AND seller_id = ?", productID, sellerID).
		Order("created_at DESC").
		Limit(limit).
		Scan(&history).Error

	if err != nil {
		return nil, err
	}
	return history, nil
}

// func (r *InventoryRepository) GetStockHistory(productID uint, sellerID uint, limit int) ([]inventory.StockHistoryResponse, error) {
// 	if limit <= 0 {
// 		limit = 50
// 	}
// 	if limit > 500 {
// 		limit = 500 // max limit
// 	}

// 	var history []inventory.StockHistoryResponse

// 	err := r.db.Raw(`
// 		SELECT
// 			sh.id, sh.product_id, sh.seller_id, sh.type, sh.quantity, sh.prev_stock, sh.new_stock,
// 			sh.reason, sh.reference, sh.created_by, sh.created_at,
// 			p.name AS product_name, p.sku AS product_sku, p.price AS product_price, p.quantity AS product_stock,
// 			c.name AS category_name,
// 			pi.image_url AS product_image
// 		FROM stock_histories sh
// 		LEFT JOIN products p ON sh.product_id = p.id
// 		LEFT JOIN categories c ON p.category_id = c.id
// 		LEFT JOIN LATERAL (
// 			SELECT image_url
// 			FROM product_images
// 			WHERE product_id = p.id
// 			ORDER BY is_primary DESC, sort_order ASC
// 			LIMIT 1
// 		) pi ON true
// 		WHERE sh.product_id = ? AND sh.seller_id = ? AND sh.deleted_at IS NULL
// 		ORDER BY sh.created_at DESC
// 		LIMIT ?
// 	`, productID, sellerID, limit).Scan(&history).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return history, nil
// }
