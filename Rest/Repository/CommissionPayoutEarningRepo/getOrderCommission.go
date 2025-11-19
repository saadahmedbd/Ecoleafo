package commissionpayoutearningrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *commissionRepository) GetOrderCommission(orderID uint) (*models.OrderCommission, error) {
	var commission models.OrderCommission
	err := r.db.Preload("Order").Preload("Seller").
		Where("order_id = ?", orderID).First(&commission).Error
	return &commission, err
}
