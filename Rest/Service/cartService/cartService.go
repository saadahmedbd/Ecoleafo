package cartservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
	cartitemrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CartItemRepo"
)

type CartService interface {
	//cart management
	AddToCart(buyerID uint, req cartitem.AddToCartRequest) (*cartitem.CartSummaryResponse, error)
	GetCart(buyerID uint) (*cartitem.CartSummaryResponse, error)
	UpdateCartItem(buyerID uint, productID uint, req cartitem.UpdateCartItemRequest) error
	RemoveFromCart(buyerID uint, productID uint) error
	ClearCart(buyerID uint) error

	// Quantity Operations
	IncrementQuantity(buyerID uint, productID uint) error
	DecrementQuantity(buyerID uint, productID uint) error
	BulkUpdateCart(buyerID uint, req cartitem.BulkUpdateCartRequest) error

	// Save for Later
	SaveForLater(buyerID uint, productID uint) error
	MoveToCart(buyerID uint, productID uint) error
	GetSavedForLater(buyerID uint) ([]models.CartItem, error)

	// Wishlist
	AddToWishlist(buyerID uint, productID uint) error
	MoveToWishlist(buyerID uint, productID uint) error
	MoveFromWishlistToCart(buyerID uint, productID uint) error
	GetWishlist(buyerID uint) (*cartitem.WishlistResponse, error)
	RemoveFromWishlist(buyerID uint, productID uint) error
	GetWishlistCount(buyerID uint) (int, error)

	// Utilities
	ValidateCart(buyerID uint) (*cartitem.CartSummaryResponse, error)
	RemoveUnavailableItems(buyerID uint) error
	GetCartSummary(buyerID uint) (*cartitem.CartSummaryResponse, error)
	GetCartItemCount(buyerID uint) (int, error)
	GetCartTotal(buyerID uint) (float64, error)
}

type cartService struct {
	cartRepo cartitemrepo.CartRepository
}

func NewCartService(cartRepo cartitemrepo.CartRepository) CartService {
	return &cartService{
		cartRepo: cartRepo,
	}
}
