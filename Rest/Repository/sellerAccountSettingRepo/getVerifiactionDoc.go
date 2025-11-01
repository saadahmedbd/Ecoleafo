package selleraccountsettingrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetVerificationDocuments retrieves all verification documents for seller
func (r *SellerAccountSettingRepository) GetVerificationDocuments(sellerID uint) ([]models.SellerVerificationDocument, error) {
	var docs []models.SellerVerificationDocument
	err := r.db.Where("seller_id = ?", sellerID).Find(&docs).Error
	return docs, err
}
