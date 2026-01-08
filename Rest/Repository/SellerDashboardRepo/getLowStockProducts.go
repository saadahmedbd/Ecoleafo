package sellerdashboardrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetLowStockProducts - Get products with low stock
func (r *DashboardRepository) GetLowStockProducts(regUserID uint, threshold int) ([]sellerdashboard.LowStockProductResponse, error) {
	// Convert reguser ID to seller ID
	var seller struct {
		ID uint
	}
	if err := r.db.Table("users").Select("id").Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return nil, err
	}
	sellerID := seller.ID

	var products []sellerdashboard.LowStockProductResponse

	err := r.db.Model(&models.Product{}).
		Select("id, name, sku, quantity, min_quantity as min_stock, price, is_active").
		Where("seller_id = ? AND quantity <= ?", sellerID, threshold).
		Order("quantity ASC").
		Scan(&products).Error

	if err != nil {
		return nil, err
	}

	// Get images and set status
	for i := range products {
		var image models.ProductImage
		r.db.Where("product_id = ?", products[i].ID).
			Order("is_primary DESC, sort_order ASC").
			First(&image)
		products[i].ImageURL = image.ImageURL

		if products[i].Quantity == 0 {
			products[i].Status = "out-of-stock"
		} else {
			products[i].Status = "low-stock"
		}
	}

	return products, nil
}
