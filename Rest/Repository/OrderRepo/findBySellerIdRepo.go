package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) FindBySellerID(sellerID uint, page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit

	subQuery := r.db.Model(&models.OrderItem{}).
		Select("DISTINCT order_id").
		Where("seller_id = ?", sellerID)

	if err := r.db.Model(&models.Order{}).Where("id IN (?)", subQuery).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Where("id IN (?)", subQuery).
		Preload("Buyer").
		Preload("OrderItems", "seller_id = ?", sellerID).
		Preload("OrderItems.Product").
		Preload("OrderItems.Product.Images").
		Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&orders).Error

	return orders, total, err
}
