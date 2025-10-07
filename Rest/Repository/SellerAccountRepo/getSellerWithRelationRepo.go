package selleraccountrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *sellerRegistrationRepository) GetSellerWithRelation(sellerID uint) (*models.User, error) {
	var Seller models.User
	err := r.db.
		Preload("RegUser").
		Preload("Role").
		Preload("PaymentMethods", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", true).Order("is_default DESC")
		}).
		Preload("Products", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,seller_id,name,price,is_active")
		}).
		Where("id = ?", sellerID).
		First(&Seller).Error
	if err != nil {
		return nil, err
	}
	return &Seller, nil
}
