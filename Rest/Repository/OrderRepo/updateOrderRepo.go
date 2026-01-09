package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (h *orderRepository) Update(order *models.Order) error {
	return h.db.Save(order).Error
}
