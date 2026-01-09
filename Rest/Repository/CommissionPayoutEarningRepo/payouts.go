package commissionpayoutearningrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// ============================================================================
// PAYOUTS
// ============================================================================
func (r *commissionRepository) CreatePayout(payout *models.SellerPayout) error {
	payout.RequestedAt = time.Now()
	return r.db.Create(payout).Error
}

func (r *commissionRepository) GetPayoutByID(id uint) (*models.SellerPayout, error) {
	var payout models.SellerPayout
	err := r.db.Preload("Seller").Preload("ProcessedByAdmin").
		First(&payout, id).Error
	return &payout, err
}

func (r *commissionRepository) GetPendingPayouts(page, limit int) ([]models.SellerPayout, int64, error) {
	var payouts []models.SellerPayout
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.SellerPayout{}).
		Where("status = ?", "pending").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Seller").
		Where("status = ?", "pending").
		Order("requested_at ASC").
		Offset(offset).Limit(limit).
		Find(&payouts).Error

	return payouts, total, err
}

func (r *commissionRepository) GetPayoutHistory(sellerID *uint, page, limit int) ([]models.SellerPayout, int64, error) {
	var payouts []models.SellerPayout
	var total int64

	offset := (page - 1) * limit

	query := r.db.Model(&models.SellerPayout{})
	if sellerID != nil {
		query = query.Where("seller_id = ?", *sellerID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Seller").Preload("ProcessedByAdmin").
		Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&payouts).Error

	return payouts, total, err
}

func (r *commissionRepository) UpdatePayoutStatus(id uint, status string, processedBy uint) error {
	updates := map[string]interface{}{
		"status":       status,
		"processed_by": processedBy,
	}

	now := time.Now()
	if status == "processing" || status == "completed" {
		updates["processed_at"] = &now
	}
	if status == "completed" {
		updates["completed_at"] = &now
	}

	return r.db.Model(&models.SellerPayout{}).Where("id = ?", id).Updates(updates).Error
}
