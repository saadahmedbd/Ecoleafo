package buyerservice

import buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"

func (s *buyerService) GetBuyerProfile(userIDFromJWT uint) (*buyerprofile.BuyerProfileResponse, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	buyerWithRelations, err := s.buyerRepo.GetBuyerWithRelations(buyer.ID)
	if err != nil {
		return nil, err
	}

	return s.mapToBuyerProfileResponse(buyerWithRelations), nil
}
