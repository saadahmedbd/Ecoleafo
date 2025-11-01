package selleraccountsettingservice

import selleraccountsettingrepo "github.com/saadahmedbd/Treestore/Rest/Repository/sellerAccountSettingRepo"

type SellerAccountSettingService struct {
	repo *selleraccountsettingrepo.SellerAccountSettingRepository
}

func NewSellerAccountSettingService(repo *selleraccountsettingrepo.SellerAccountSettingRepository) *SellerAccountSettingService {
	return &SellerAccountSettingService{repo: repo}
}
