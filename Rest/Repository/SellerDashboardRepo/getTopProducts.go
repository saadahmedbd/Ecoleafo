package sellerdashboardrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetTopProducts - Get best selling products
func (r *DashboardRepository) GetTopProducts(regUserID uint, limit int) ([]sellerdashboard.TopProductResponse, error) {
	// Convert reguser ID to seller ID
	var seller struct {
		ID uint
	}
	if err := r.db.Table("users").Select("id").Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return nil, err
	}
	sellerID := seller.ID

	var products []sellerdashboard.TopProductResponse

	err := r.db.Table("products").
		Select(`products.id, products.name, products.sku, products.price, 
			COALESCE(SUM(order_items.quantity), 0) as total_sold,
			COALESCE(SUM(order_items.seller_earning), 0) as revenue,
			products.quantity as stock, products.is_active, products.is_approved`).
		Joins("LEFT JOIN order_items ON products.id = order_items.product_id").
		Where("products.seller_id = ?", sellerID).
		Group("products.id").
		Order("total_sold DESC").
		Limit(limit).
		Scan(&products).Error

	if err != nil {
		return nil, err
	}

	// Get first image for each product
	for i := range products {
		var image models.ProductImage
		r.db.Where("product_id = ?", products[i].ID).
			Order("is_primary DESC, sort_order ASC").
			First(&image)
		products[i].ImageURL = image.ImageURL
	}

	return products, nil
}
