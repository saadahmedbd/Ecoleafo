package buyerservice

import (
	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
	buyerProfilerepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerProfileRepo"
)

type BuyerService interface {
	GetBuyerProfile(userIDFromJWT uint) (*buyerprofile.BuyerProfileResponse, error)
	UpdateBuyerProfile(userIDFromJWT uint, req buyerprofile.UpdateBuyerProfileRequest) (*buyerprofile.BuyerProfileResponse, error)
	ChangePassword(userIDFromJWT uint, req buyerprofile.ChangeBuyerPassword) error
	GetBuyerStats(userIDFromJWT uint) (*buyerprofile.BuyerStateResponse, error)
	GetBuyerOrderHistory(userIDFromJWT uint, page, limit int) (*buyerprofile.BuyerOrderHistoryResponse, error)

	// Address management
	CreateAddress(userIDFromJWT uint, req buyerprofile.AddressInfo) (*buyerprofile.AddressInfo, error)
	UpdateAddress(userIDFromJWT uint, addressID uint, req buyerprofile.UpdateBuyerAddressRequest) (*buyerprofile.AddressInfo, error)
	DeleteAddress(userIDFromJWT uint, addressID uint) error
	GetAddresses(userIDFromJWT uint) ([]buyerprofile.AddressInfo, error)
	SetDefaultAddress(userIDFromJWT uint, addressID uint) error
	UpdateProfilePhoto(buyerID uint, photoURL, publicID string) (*buyerprofile.UpdateProfilePhotoResponse, error)
	DeleteBuyerProfilePicture(buyerID uint, imageURL string) error
}

type buyerService struct {
	buyerRepo buyerProfilerepo.BuyerRepository
}

func NewBuyerService(buyerRepo buyerProfilerepo.BuyerRepository) BuyerService {
	return &buyerService{
		buyerRepo: buyerRepo,
	}
}
