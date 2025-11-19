package commissionpayoutearningrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *commissionRepository) UpdateCommissionSettings(settings *models.CommissionSetting) error {
	return r.db.Save(settings).Error
}
