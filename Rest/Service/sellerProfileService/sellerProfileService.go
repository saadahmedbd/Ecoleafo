package sellerprofileservice

import (
	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
	repository "github.com/saadahmedbd/Treestore/Rest/Repository"
)

type SellerService interface {
	GetSellerProfile(userIDFromJWT uint) (*sellerprofile.SellerProfileResponse, error)
	UpdateSellerProfile(userIDFromJWT uint, req sellerprofile.UpdateSellerProfileRequest) (*sellerprofile.SellerProfileResponse, error)
	UpdateStoreImages(userIDFromJWT uint, req sellerprofile.UpdateStoreImageRequest) error
	ChangePassword(userIDFromJWT uint, req sellerprofile.ChangePasswordRequest) error
	GetSellerStats(userIDFromJWT uint) (*sellerprofile.SellerStateResponse, error)
	GetPublicSellerProfile(sellerID uint) (*sellerprofile.SellerProfileResponse, error)
}

type sellerService struct {
	sellerRepo repository.SellerRepository
}

func NewSellerService(sellerRepo repository.SellerRepository) SellerService {
	return &sellerService{
		sellerRepo: sellerRepo,
	}
}
