package buyerservice

import buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"

func (s *buyerService) GetBuyerStats(userIDFromJWT uint) (*buyerprofile.BuyerStateResponse, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	return s.buyerRepo.GetBuyerStats(buyer.ID)
}
