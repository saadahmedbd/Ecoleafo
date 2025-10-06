package guestcartservice

import (
	guestcartitem "github.com/saadahmedbd/Treestore/Rest/DTO/GuestCartItem"
)

func (s *guestcartservice) GetCart(sessionID string, userID *uint) (*guestcartitem.CartResponse, error) {
	response := &guestcartitem.CartResponse{
		Items:       []guestcartitem.CartItemInfo{},
		IsGuest:     userID == nil,
		CanCheckout: false,
	}

	if userID != nil {
		// Registered user
		buyer, err := s.buyerRepo.GetBuyerByUserID(*userID)
		if err != nil {
			return nil, err
		}

		items, err := s.cartRepo.GetUserCart(buyer.ID)
		if err != nil {
			return nil, err
		}

		for _, item := range items {
			imageURL := ""
			if len(item.Product.Images) > 0 {
				imageURL = item.Product.Images[0].ImageURL
			}

			response.Items = append(response.Items, guestcartitem.CartItemInfo{
				ID:          item.ID,
				ProductID:   item.ProductID,
				ProductName: item.Product.Name,
				ProductSlug: item.Product.Slug,
				Price:       item.Price,
				Quantity:    item.Quantity,
				Subtotal:    item.Price * float64(item.Quantity),
				Image:       imageURL,
				InStock:     item.Product.Quantity >= item.Quantity,
			})

			response.SubTotal += item.Price * float64(item.Quantity)
		}

		response.TotalItems = len(items)

		// Check if profile is complete
		isComplete, _, _ := s.profileRepo.HasCompletedProfile(buyer.ID)
		response.CanCheckout = isComplete
		response.RequiresProfile = !isComplete
	} else {
		// Guest user
		if sessionID == "" {
			return response, nil
		}

		response.SessionID = sessionID

		items, err := s.cartRepo.GetGuestCart(sessionID)
		if err != nil {
			return nil, err
		}

		for _, item := range items {
			imageURL := ""
			if len(item.Product.Images) > 0 {
				imageURL = item.Product.Images[0].ImageURL
			}

			response.Items = append(response.Items, guestcartitem.CartItemInfo{
				ID:          item.ID,
				ProductID:   item.ProductID,
				ProductName: item.Product.Name,
				ProductSlug: item.Product.Slug,
				Price:       item.Price,
				Quantity:    item.Quantity,
				Subtotal:    item.Price * float64(item.Quantity),
				Image:       imageURL,
				InStock:     item.Product.Quantity >= item.Quantity,
			})

			response.SubTotal += item.Price * float64(item.Quantity)
		}

		response.TotalItems = len(items)
		response.CanCheckout = false
		response.RequiresProfile = true
	}

	return response, nil
}
