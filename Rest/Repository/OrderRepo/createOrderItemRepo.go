package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) CreateOrderItem(item *models.OrderItem) error {
	return r.db.Create(item).Error
}
