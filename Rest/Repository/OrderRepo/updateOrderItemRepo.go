package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) UpdateOrderItem(item *models.OrderItem) error {
	return r.db.Save(item).Error
}
