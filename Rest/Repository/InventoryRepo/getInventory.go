package inventoryrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
)

func (r *InventoryRepository) GetInventoryItem(productID uint) (*inventory.InventoryItemResponse, error) {
	var item inventory.InventoryItemResponse

	err := r.db.Table("products").
		Select(`products.id, products.name, products.sku, products.category_id, 
			categories.name as category_name, products.price, products.quantity as stock, 
			products.min_quantity as low_stock_threshold, products.is_active, products.is_approved,
			products.updated_at as last_updated,
			(products.quantity * products.price) as stock_value`).
		Joins("LEFT JOIN categories ON products.category_id = categories.id").
		Where("products.id = ?", productID).
		Scan(&item).Error

	if err != nil {
		return nil, err
	}

	var image models.ProductImage
	r.db.Where("product_id = ?", productID).
		Order("is_primary DESC, sort_order ASC").
		First(&image)
	item.ImageURL = image.ImageURL

	return &item, nil
}

func (r *InventoryRepository) GetInventory(sellerID uint, filters inventory.InventoryFilter) (*inventory.InventoryListResponse, error) {
	var products []models.Product
	var total int64

	// Debug: Check what products exist for this seller
	var count int64
	r.db.Model(&models.Product{}).Where("seller_id = ?", sellerID).Count(&count)
	if count == 0 {
		// No products found, return empty list
		return &inventory.InventoryListResponse{
			Items: []inventory.InventoryItemResponse{},
			Total: 0,
			Page:  filters.Page,
			Limit: filters.Limit,
		}, nil
	}

	query := r.db.Model(&models.Product{}).Where("seller_id = ?", sellerID)

	if filters.Search != "" {
		searchTerm := "%" + filters.Search + "%"
		query = query.Where("name LIKE ? OR sku LIKE ?", searchTerm, searchTerm)
	}

	if filters.Status != "" {
		switch filters.Status {
		case "low-stock":
			query = query.Where("quantity > 0 AND quantity <= min_quantity")
		case "out-of-stock":
			query = query.Where("quantity = 0")
		case "in-stock":
			query = query.Where("quantity > min_quantity")
		}
	}

	if filters.Category != "" {
		query = query.Where("category_id = ?", filters.Category)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (filters.Page - 1) * filters.Limit
	err := query.Preload("Category").Order("updated_at DESC").
		Offset(offset).Limit(filters.Limit).
		Find(&products).Error
	if err != nil {
		return nil, err
	}

	items := make([]inventory.InventoryItemResponse, len(products))
	for i, p := range products {
		items[i] = inventory.InventoryItemResponse{
			ID:                p.ID,
			Name:              p.Name,
			SKU:               p.SKU,
			CategoryID:        p.CategoryID,
			CategoryName:      p.Category.Name,
			Price:             p.Price,
			Stock:             p.Quantity,
			LowStockThreshold: p.MinQuantity,
			IsActive:          p.IsActive,
			IsApproved:        p.IsApproved,
			LastUpdated:       p.UpdatedAt,
			StockValue:        float64(p.Quantity) * p.Price,
		}

		var image models.ProductImage
		if err := r.db.Where("product_id = ?", p.ID).
			Order("is_primary DESC, sort_order ASC").
			First(&image).Error; err == nil {
			items[i].ImageURL = image.ImageURL
		}
	}

	return &inventory.InventoryListResponse{
		Items: items,
		Total: total,
		Page:  filters.Page,
		Limit: filters.Limit,
	}, nil
}
