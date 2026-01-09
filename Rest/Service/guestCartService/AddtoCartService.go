package guestcartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	guestcartitem "github.com/saadahmedbd/Treestore/Rest/DTO/GuestCartItem"
)

// get product and verify price and stock
func (s *guestcartservice) AddToCart(sessionID string, userID *uint, req guestcartitem.AddToCartRequest) (*guestcartitem.CartResponse, error) {
	// Get product to verify price and stock
	var product models.Product
	if err := Config.DB.Where("id = ?", req.ProductID).First(&product).Error; err != nil {
		return nil, fmt.Errorf("product not found")
	}

	if product.Quantity < req.Quantity {
		return nil, fmt.Errorf("insufficient stock")
	}

	if userID != nil {
		// Registered user
		buyer, err := s.buyerRepo.GetBuyerByUserID(*userID)
		if err != nil {
			return nil, err
		}

		cartItem := &models.CartItem{
			BuyerID:   buyer.ID,
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
			Price:     product.Price,
		}

		if err := s.cartRepo.AddCartItem(cartItem); err != nil {
			return nil, err
		}
	} else {
		// Guest user
		if sessionID == "" {
			return nil, fmt.Errorf("session ID required for guest users")
		}

		guestItem := &models.GuestCartItem{
			SessionID: sessionID,
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
			Price:     product.Price,
		}

		if err := s.cartRepo.AddGuestCartItem(guestItem); err != nil {
			return nil, err
		}
	}

	return s.GetCart(sessionID, userID)
}
