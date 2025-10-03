package buyerservice

import (
	"fmt"
	models "github.com/saadahmedbd/Treestore/Models"
	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
)

// Address management methods
func (s *buyerService) CreateAddress(userIDFromJWT uint, req buyerprofile.AddressInfo) (*buyerprofile.AddressInfo, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	address := &models.Address{
		BuyerID:      buyer.ID,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		Street:       req.Street,
		City:         req.City,
		State:        req.State,
		District:     req.District,
		Country:      req.Country,
		PostalCode:   req.PostalCode,
		IsDefault:    req.IsDefault,
	}

	if err := s.buyerRepo.CreateAddress(address); err != nil {
		fmt.Printf("DEBUG - Error creating address: %v\n", err)
		return nil, err
	}

	if req.IsDefault {
		s.buyerRepo.SetDefaultAddress(buyer.ID, address.ID)
	}

	return &buyerprofile.AddressInfo{
		
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
