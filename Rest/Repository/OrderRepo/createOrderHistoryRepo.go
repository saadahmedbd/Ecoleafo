package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) CreateOrderHistory(history *models.OrderHistory) error {
	return r.db.Create(history).Error
}
