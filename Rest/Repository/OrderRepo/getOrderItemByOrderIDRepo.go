package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) GetOrderItemByOrderID(orderID uint) ([]models.OrderItem, error) {
	var items []models.OrderItem
	err := r.db.Preload("Product").
		Preload("Seller").
		Where("order_id = ?", orderID).Find(&items).Error

	return items, err

}
