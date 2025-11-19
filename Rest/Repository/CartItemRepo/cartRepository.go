package cartitemrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type CartRepository interface {
	//cart operation
	AddToCart(item *models.CartItem) error
	GetCart(buyerID uint, includeSavedForLater bool) ([]models.CartItem, error)
	GetCartItem(buyerID, productID uint) (*models.CartItem, error)
	UpdateCartItem(item *models.CartItem) error
	RemoveFromCart(buyerID, productID uint) error
	ClearCart(buyerID uint) error

	// Quantity Management
	UpdateQuantity(buyerID, productID uint, quantity int) error
	IncrementQuantity(buyerID, productID uint) error
	DecrementQuantity(buyerID, productID uint) error

	// Save for Later
	SaveForLater(buyerID, productID uint) error
	MoveToCart(buyerID, productID uint) error
	GetSavedForLater(buyerID uint) ([]models.CartItem, error)

	// Wishlist
	AddToWishlist(wishlist *models.Wishlist) error
	MoveToWishlist(buyerID, productID uint) error
	MoveFromWishlistToCart(buyerID, productID uint, quantity int) error
	GetWishlist(buyerID uint) ([]models.Wishlist, error)
	RemoveFromWishlist(buyerID, productID uint) error
	GetWishlistCount(buyerID uint) (int, error)

	// Bulk Operations
	BulkUpdateQuantities(buyerID uint, updates map[uint]int) error
	RemoveUnavailableItems(buyerID uint) error

	// Validation
	ValidateCartStock(buyerID uint) ([]uint, error) // Returns product IDs with stock issues
	GetCartItemCount(buyerID uint) (int, error)
	GetCartTotal(buyerID uint) (float64, error)
}

type cartRepository struct {
	// Add any dependencies or configurations here
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}
