package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) FindAll(filter map[string]interface{}, page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64
	offset := (page - 1) * limit
	query := r.db.Model(&models.Order{})
	for key, value := range filter {
		query = query.Where(key+"= ?", value)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Buyer").
		Preload("OrderItems.Product").
		Preload("OrderItems.Seller").
		Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&orders).Error

	return orders, total, err
}
