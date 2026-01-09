package selleraccountsettingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

// UpdateVerificationDocumentStatus updates document status
func (r *SellerAccountSettingRepository) UpdateVerificationDocumentStatus(docID uint, status string) error {
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	if status == "verified" {
		now := time.Now()
		updates["verified_at"] = &now
	} else if status == "rejected" {
		now := time.Now()
		updates["rejected_at"] = &now
	}

	return r.db.Model(&models.SellerVerificationDocument{}).
		Where("id = ?", docID).
		Updates(updates).Error
}
