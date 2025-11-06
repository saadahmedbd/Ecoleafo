package selleraccountsettingservice

import selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"

// GetStatistics retrieves seller statistics
func (s *SellerAccountSettingService) GetStatistics(userID uint) (*selleraccountsetting.SellerStatisticsResponse, error) {
	// Get seller by RegUser ID (from JWT)
	seller, err := s.repo.GetSellerByUserID(userID)
	if err != nil {
		return nil, err
	}
	// Get statistics using actual seller ID
	return s.repo.GetSellerStatisticsStruct(seller.ID)
}
