package commissionpayoutearningrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *commissionRepository) UpdateCommissionStatus(orderID uint, status string) error {
	updates := map[string]interface{}{"status": status}
	if status == "cleared" {
		now := time.Now()
		updates["cleared_at"] = &now
	} else if status == "paid_out" {
		now := time.Now()
		updates["paid_out_at"] = &now
	}
	return r.db.Model(&models.OrderCommission{}).
		Where("order_id = ?", orderID).
		Updates(updates).Error
}
