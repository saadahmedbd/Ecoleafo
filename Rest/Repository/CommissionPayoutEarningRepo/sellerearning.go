package commissionpayoutearningrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// ============================================================================
// SELLER EARNINGS
// ============================================================================
func (r *commissionRepository) GetSellerEarnings(sellerID uint) (*models.SellerEarningsSummary, error) {
	var earnings models.SellerEarningsSummary
	err := r.db.Preload("Seller").Where("seller_id = ?", sellerID).First(&earnings).Error
	if err == gorm.ErrRecordNotFound {
		// Create new earnings summary
		earnings = models.SellerEarningsSummary{
			SellerID:    sellerID,
			LastUpdated: time.Now(),
		}
		if err := r.db.Create(&earnings).Error; err != nil {
			return nil, err
		}
	}
	return &earnings, err
}

func (r *commissionRepository) UpdateSellerEarnings(earnings *models.SellerEarningsSummary) error {
	earnings.LastUpdated = time.Now()
	return r.db.Save(earnings).Error
}

func (r *commissionRepository) GetAllSellerEarnings(page, limit int) ([]models.SellerEarningsSummary, int64, error) {
	var earnings []models.SellerEarningsSummary
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.SellerEarningsSummary{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Seller").
		Order("net_earnings DESC").
		Offset(offset).Limit(limit).
		Find(&earnings).Error

	return earnings, total, err
}
