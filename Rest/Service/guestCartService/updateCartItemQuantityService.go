package guestcartservice

import (
	"fmt"
	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (s *guestcartservice) UpdateCartItemQuantity(sessionID string, userID *uint, productID uint, quantity int) error {
	if quantity <= 0 {
		return s.RemoveCartItem(sessionID, userID, productID)
	}

	if userID != nil {
		buyer, err := s.buyerRepo.GetBuyerByUserID(*userID)
		if err != nil {
			// If buyer doesn't exist, treat as guest
			if err.Error() == "buyer not found" {
				var item models.GuestCartItem
				if err := Config.DB.Where("session_id = ? AND product_id = ?", sessionID, productID).First(&item).Error; err != nil {
					return fmt.Errorf("cart item not found")
				}
				item.Quantity = quantity
				return s.cartRepo.UpdateGuestCartItem(&item)
			}
			return err
		}

		var item models.CartItem
		if err := Config.DB.Where("buyer_id = ? AND product_id = ?", buyer.ID, productID).First(&item).Error; err != nil {
			return fmt.Errorf("cart item not found")
		}

		item.Quantity = quantity
		return s.cartRepo.UpdateCartItem(&item)
	} else {
		var item models.GuestCartItem
		if err := Config.DB.Where("session_id = ? AND product_id = ?", sessionID, productID).First(&item).Error; err != nil {
			return fmt.Errorf("cart item not found")
		}

		item.Quantity = quantity
		return s.cartRepo.UpdateGuestCartItem(&item)
	}
}
