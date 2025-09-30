package sellerprofileservice

import sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"

func (s *sellerService) UpdateStoreImages(userIDFromJWT uint, req sellerprofile.UpdateStoreImageRequest) error {
	seller, err := s.sellerRepo.GetSellerByUserID(userIDFromJWT)
	if err != nil {
		return err
	}

	if req.StoreLogo != "" {
		seller.StoreLogo = req.StoreLogo
	}
	if req.StoreBanner != "" {
		seller.StoreBanner = req.StoreBanner
	}

	return s.sellerRepo.UpdateSellerProfile(seller)
}
