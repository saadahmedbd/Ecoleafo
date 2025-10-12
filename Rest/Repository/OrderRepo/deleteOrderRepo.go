package orderrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *orderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}
