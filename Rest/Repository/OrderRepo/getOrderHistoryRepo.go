package orderrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *orderRepository) GetOrderHistory(orderID uint) ([]models.OrderHistory, error) {
	var history []models.OrderHistory
	err := r.db.Where("order_id = ?", orderID).
		Order("created_at DESC").Find(&history).Error
	return history, err
}
