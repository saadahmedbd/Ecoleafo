package sellerprofileservice

import sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"

func (s *sellerService) GetSellerStats(userIDFromJWT uint) (*sellerprofile.SellerStateResponse, error) {
	seller, err := s.sellerRepo.GetSellerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	stats, err := s.sellerRepo.GetSellerStats(seller.ID)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
