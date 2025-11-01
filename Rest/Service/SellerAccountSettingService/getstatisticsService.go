package selleraccountsettingservice

import selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"

// Service becomes trivial!
func (s *SellerAccountSettingService) GetStatistics(sellerID uint) (*selleraccountsetting.SellerStatisticsResponse, error) {
	// That's it! Repository handles everything
	return s.repo.GetSellerStatisticsStruct(sellerID)
}
