package commissionpayoutearningrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *commissionRepository) CreateOrderCommission(commission *models.OrderCommission) error {
	return r.db.Create(commission).Error
}
