package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// CreateVerificationDocument creates new verification document
func (r *SellerAccountSettingRepository) CreateVerificationDocument(doc *models.SellerVerificationDocument) error {
	return r.db.Create(doc).Error
}
