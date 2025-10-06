package guestcartservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	address "github.com/saadahmedbd/Treestore/Rest/DTO/Address"
)

func (s *guestcartservice) CreateCheckoutAddress(userID uint, req address.CreateAddressRequest) (*models.Address, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userID)
	if err != nil {
		return nil, err
	}

	address := &models.Address{
		BuyerID:      buyer.ID,
		Type:         req.Type,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Company:      req.Company,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		Street:       req.Street,
		City:         req.City,
		State:        req.State,
		District:     req.District,
		PostalCode:   req.PostalCode,
		Country:      req.Country,
		Phone:        req.Phone,
		IsDefault:    req.IsDefault,
	}

	if address.Country == "" {
		address.Country = "Bangladesh"
	}

	if err := s.profileRepo.CreateAddress(address); err != nil {
		return nil, err
	}

	// Set as default address ID if it's the first address or marked as default
	if req.IsDefault {
		buyer.DefaultAddressID = &address.ID
		s.profileRepo.UpdateBuyerProfile(buyer)
	}

	return address, nil
}
