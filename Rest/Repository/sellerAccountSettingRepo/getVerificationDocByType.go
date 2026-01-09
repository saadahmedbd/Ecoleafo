package selleraccountsettingrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GetVerificationDocumentByType retrieves specific document type
func (r *SellerAccountSettingRepository) GetVerificationDocumentByType(sellerID uint, docType string) (*models.SellerVerificationDocument, error) {
	var doc models.SellerVerificationDocument
	err := r.db.Where("seller_id = ? AND document_type = ?", sellerID, docType).
		Order("created_at DESC").
		First(&doc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No document found is not an error
		}
		return nil, err
	}
	return &doc, nil
}
