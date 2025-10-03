package buyerservice

import (
	"fmt"
	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
)

func (s *buyerService) UpdateAddress(userIDFromJWT uint, addressID uint, req buyerprofile.UpdateBuyerAddressRequest) (*buyerprofile.AddressInfo, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	address, err := s.buyerRepo.GetAddressByID(addressID, buyer.ID)
	if err != nil {
		return nil, err
	}

	// Update only non-empty fields
	if req.AddressLine1 != "" {
		address.AddressLine1 = req.AddressLine1
	}
	if req.AddressLine2 != "" {
		address.AddressLine2 = req.AddressLine2
	}
	if req.Street != "" {
		address.Street = req.Street
	}
	if req.City != "" {
		address.City = req.City
	}
	if req.State != "" {
		address.State = req.State
	}
	if req.District != "" {
		address.District = req.District
	}
	if req.Country != "" {
		address.Country = req.Country
	}
	if req.PostalCode != "" {
		address.PostalCode = req.PostalCode
	}

	if err := s.buyerRepo.UpdateAddress(address); err != nil {
		fmt.Printf("DEBUG - Error updating address: %v\n", err)
		return nil, err
	}

	if req.IsDefault && !address.IsDefault {
		address.IsDefault = true
		s.buyerRepo.SetDefaultAddress(buyer.ID, address.ID)
	}

	return &buyerprofile.AddressInfo{
		ID:           address.ID,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		Street:       address.Street,
		City:         address.City,
		State:        address.State,
		District:     address.District,
		Country:      address.Country,
		PostalCode:   address.PostalCode,
		IsDefault:    address.IsDefault,
	}, nil
}
