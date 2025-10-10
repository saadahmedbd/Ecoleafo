package cartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
)

func (s *cartService) AddToCart(buyerID uint, req cartitem.AddToCartRequest) (*cartitem.CartSummaryResponse, error) {
	// Get product
	var product models.Product
	if err := Config.DB.Where("id = ?", req.ProductID).First(&product).Error; err != nil {
		return nil, fmt.Errorf("product not found")
	}

	// Validate product availability
	if !product.IsActive {
		return nil, fmt.Errorf("product is not available")
	}

	if product.Quantity < req.Quantity {
		return nil, fmt.Errorf("insufficient stock. Only %d available", product.Quantity)
	}

	// Create cart item
	cartItem := &models.CartItem{
		BuyerID:     buyerID,
		ProductID:   req.ProductID,
		Quantity:    req.Quantity,
		Price:       product.Price,
		IsGift:      req.IsGift,
		GiftMessage: req.GiftMessage,
		AddedFrom:   "web",
	}

	if err := s.cartRepo.AddToCart(cartItem); err != nil {
		return nil, err
	}

	return s.GetCart(buyerID)
}
