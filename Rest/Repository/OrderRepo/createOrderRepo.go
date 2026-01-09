package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}
