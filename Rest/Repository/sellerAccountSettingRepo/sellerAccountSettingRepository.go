package selleraccountsettingrepo

import "gorm.io/gorm"

type SellerAccountSettingRepository struct {
	db *gorm.DB
}

func NewSellerAccountSettingRepositoryOptions(db *gorm.DB) *SellerAccountSettingRepository {
	return &SellerAccountSettingRepository{
		db: db,
	}
}
