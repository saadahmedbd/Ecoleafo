package guestcartservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	address "github.com/saadahmedbd/Treestore/Rest/DTO/Address"
	buyeraccount "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerAccount"
	guestcartitem "github.com/saadahmedbd/Treestore/Rest/DTO/GuestCartItem"
	buyercompleterepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerCompleteRepo"
	buyerProfilerepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerProfileRepo"
	guestcartrepo "github.com/saadahmedbd/Treestore/Rest/Repository/GuestCartRepo"
)

type Guestcartservice interface {
	//cart operation
	AddToCart(sessionID string, userID *uint, req guestcartitem.AddToCartRequest) (*guestcartitem.CartResponse, error)
	GetCart(sessionID string, userID *uint) (*guestcartitem.CartResponse, error)
	UpdateCartItemQuantity(sessionID string, userID *uint, productID uint, quantity int) error
	RemoveCartItem(sessionID string, userID *uint, productID uint) error

	// Profile completion
	CheckProfileCompletion(userID uint) (*buyeraccount.ProfileCompletionStatus, error)
	CompleteProfile(userID uint, req buyeraccount.CompleteProfileRequest) error
	CreateCheckoutAddress(userID uint, req address.CreateAddressRequest) (*models.Address, error)

	// checkout
	InitiateCheckout(userID uint, req address.CheckoutInitRequest) error
	MigrateGuestCart(sessionID string, userID uint) error
}
type guestcartservice struct {
	cartRepo    guestcartrepo.GuestCartRepository
	profileRepo buyercompleterepo.ProfileCompleteRepository
	buyerRepo   buyerProfilerepo.BuyerRepository
}

func NewGuestCartService(cartRepo guestcartrepo.GuestCartRepository,
	profileRepo buyercompleterepo.ProfileCompleteRepository,
	buyerRepo buyerProfilerepo.BuyerRepository) Guestcartservice {
	return &guestcartservice{
		cartRepo:    cartRepo,
		profileRepo: profileRepo,
		buyerRepo:   buyerRepo,
	}
}
