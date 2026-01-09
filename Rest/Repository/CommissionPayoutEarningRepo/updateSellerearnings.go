package commissionpayoutearningrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *commissionRepository) UpdateSellerEarning(earnings *models.SellerEarningsSummary) error {
	earnings.LastUpdated = time.Now()
	return r.db.Save(earnings).Error
}
