package buyerservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
)

func (s *buyerService) mapToBuyerProfileResponse(buyer *models.Buyer) *buyerprofile.BuyerProfileResponse {
	response := &buyerprofile.BuyerProfileResponse{
		ID:               buyer.ID,
		UserId:           buyer.UserId,
		Phone:            buyer.Phone,
		DefaultAddress:   buyer.DefaultAddress,
		IsActive:         buyer.IsActive,
		EmailVerified:    buyer.EmailVerified,
		LastOrderAt:      buyer.LastOrderAt,
		TotalOrdersCount: buyer.TotalOrdersCount,
		TotalSpent:       buyer.TotalSpent,
		CreatedAt:        buyer.CreatedAt,
		UpdatedAt:        buyer.UpdatedAt,
		RegUser: sellerprofile.RegUserBasicInfo{
			FirstName: buyer.RegUser.FirstName,
			LastName:  buyer.RegUser.LastName,
			Email:     buyer.RegUser.Email,
		},
		WishlistCount: len(buyer.Wishlists),
		CartItemCount: len(buyer.CartItems),
	}

	if buyer.DefaultAddr != nil {
		response.DefaultAddr = &buyerprofile.AddressInfo{
			ID:           buyer.DefaultAddr.ID,
			AddressLine1: buyer.DefaultAddr.AddressLine1,
			AddressLine2: buyer.DefaultAddr.AddressLine2,
			City:         buyer.DefaultAddr.City,
			State:        buyer.DefaultAddr.State,
			Country:      buyer.DefaultAddr.Country,
			PostalCode:   buyer.DefaultAddr.PostalCode,
			IsDefault:    buyer.DefaultAddr.IsDefault,
		}
	}

	for _, addr := range buyer.Addresses {
		response.Addresses = append(response.Addresses, buyerprofile.AddressInfo{
			ID:           addr.ID,
			AddressLine1: addr.AddressLine1,
			AddressLine2: addr.AddressLine2,
			City:         addr.City,
			State:        addr.State,
			Country:      addr.Country,
			PostalCode:   addr.PostalCode,
			IsDefault:    addr.IsDefault,
		})
	}

	return response
}
