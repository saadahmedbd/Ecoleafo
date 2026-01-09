package sellerprofileservice

import sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"

func (s *sellerService) GetSellerProfile(userIDFromJWT uint) (*sellerprofile.SellerProfileResponse, error) {
	// Get seller by user_id from JWT
	seller, err := s.sellerRepo.GetSellerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	// Get seller with relations
	sellerWithRelations, err := s.sellerRepo.GetSellerWithRelations(seller.ID)
	if err != nil {
		return nil, err
	}

	return s.mapToSellerProfileResponse(sellerWithRelations), nil
}
