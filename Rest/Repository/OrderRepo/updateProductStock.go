package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) UpdateStock(id uint, newStock int) error {
	return r.db.Model(&models.Product{}).
		Where("id = ?", id).
		Update("quantity", newStock).Error
}
