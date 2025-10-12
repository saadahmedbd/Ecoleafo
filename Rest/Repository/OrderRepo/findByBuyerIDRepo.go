package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) FindByBuyerID(buyerID uint, page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit
	if err := r.db.Model(&models.Order{}).
		Where("buyer_id = ?", buyerID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.Where("buyer_id = ?", buyerID).
		Preload("OrderItems.Product").
		Preload("OrderItems.Seller").
		Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&orders).Error

	return orders, total, err
}
