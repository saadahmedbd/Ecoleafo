package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

//for updating stock(i was not create product stock that why i need to create that)
func (r *orderRepository) UpdateStock(id uint, newStock int) error {
	return r.db.Model(&models.Product{}).
		Where("id = ?", id).
		Update("quantity", newStock).Error
}
