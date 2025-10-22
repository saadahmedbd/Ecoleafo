package buyerservice

import (
	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
)

func (s *buyerService) GetAddresses(userIDFromJWT uint) ([]buyerprofile.AddressInfo, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	addresses, err := s.buyerRepo.GetBuyerAddresses(buyer.ID)
	if err != nil {
		return nil, err
	}

	var addressInfos []buyerprofile.AddressInfo
	for _, addr := range addresses {
		addressInfos = append(addressInfos, buyerprofile.AddressInfo{
			ID:           addr.ID,
			FullName:     addr.FullName,
			AddressLine1: addr.AddressLine1,
			AddressLine2: addr.AddressLine2,
			City:         addr.City,
			Street:       addr.Street,
			District:     addr.District,
			State:        addr.State,
			Country:      addr.Country,
			Phone:        addr.Phone,
			PostalCode:   addr.PostalCode,
			IsDefault:    addr.IsDefault,
		})
	}

	return addressInfos, nil
}

func (s *buyerService) SetDefaultAddress(userIDFromJWT uint, addressID uint) error {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return err
	}

	return s.buyerRepo.SetDefaultAddress(buyer.ID, addressID)
}
