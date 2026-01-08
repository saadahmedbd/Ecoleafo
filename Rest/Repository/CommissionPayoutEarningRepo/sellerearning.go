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
	err := r.db.Preload("Seller.RegUser").Where("seller_id = ?", sellerID).First(&earnings).Error
	if err == gorm.ErrRecordNotFound {
		// Create new earnings summary
		earnings = models.SellerEarningsSummary{
			SellerID:    sellerID,
			LastUpdated: time.Now(),
		}
		if err := r.db.Create(&earnings).Error; err != nil {
			return nil, err
		}
		// Reload with preloads
		if err := r.db.Preload("Seller.RegUser").Where("seller_id = ?", sellerID).First(&earnings).Error; err != nil {
			return nil, err
		}
		return &earnings, nil
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

	// Count all sellers instead of just earnings summaries
	if err := r.db.Model(&models.User{}).Where("role_id = ?", 2).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get all sellers with their earnings summaries (left join)
	var sellers []models.User
	err := r.db.Preload("RegUser").
		Where("role_id = ?", 2).
		Order("total_sales DESC").
		Offset(offset).Limit(limit).
		Find(&sellers).Error
	
	if err != nil {
		return nil, 0, err
	}
	
	// Create earnings summary for each seller
	for _, seller := range sellers {
		var summary models.SellerEarningsSummary
		err := r.db.Where("seller_id = ?", seller.ID).First(&summary).Error
		if err != nil {
			// Create new summary if not exists
			summary = models.SellerEarningsSummary{
				SellerID: seller.ID,
				Seller:   seller,
			}
		} else {
			summary.Seller = seller
		}
		earnings = append(earnings, summary)
	}

	return earnings, total, nil
}
