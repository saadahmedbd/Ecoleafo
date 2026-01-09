package commissionpayoutearningrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *commissionRepository) GetCommissionSettings() (*models.CommissionSetting, error) {
	var settings models.CommissionSetting
	err := r.db.Where("is_active = ?", true).First(&settings).Error
	if err == gorm.ErrRecordNotFound {
		// Create default settings if none exist
		settings = models.CommissionSetting{
			DefaultRate: 10.00,
			Description: "Default platform commission",
			IsActive:    true,
			UpdatedBy:   4,
		}
		if err := r.db.Create(&settings).Error; err != nil {
			return nil, err
		}
	}
	return &settings, err
}
