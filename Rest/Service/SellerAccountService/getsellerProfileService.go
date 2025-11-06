package selleraccountservice

import selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"

func (s *sellerRegistrationService) GetSellerProfile(userID uint) (*selleraccount.SellerProfileResponse, error) {
	seller, err := s.sellerRepo.GetSellerByRegUserId(userID)
	if err != nil {
		return nil, err
	}

	sellerWithRelations, err := s.sellerRepo.GetSellerWithRelation(seller.ID)
	if err != nil {
		return nil, err
	}

	return s.mapToSellerProfileResponse(sellerWithRelations), nil
}
