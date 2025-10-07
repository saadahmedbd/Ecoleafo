package selleraccountservice

import (
	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *sellerRegistrationService) UpdateStoreInfo(userID uint, req *selleraccount.UpdateStoreInfoRequest) error {
	seller, err := s.sellerRepo.GetSellerByUserId(userID)
	if err != nil {
		return err
	}

	if req.StoreName != "" {
		seller.StoreName = req.StoreName
		seller.StoreSlug = util.GenerateSlug(req.StoreName)
	}
	if req.StoreDesc != "" {
		seller.StoreDesc = req.StoreDesc
	}
	if req.StoreLogo != "" {
		seller.StoreLogo = req.StoreLogo
	}
	if req.StoreBanner != "" {
		seller.StoreBanner = req.StoreBanner
	}

	return s.sellerRepo.UpdateSellerProfile(seller)
}
