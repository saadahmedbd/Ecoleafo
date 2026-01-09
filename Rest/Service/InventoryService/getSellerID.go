package inventoryservice

import "errors"

func (s *InventoryService) GetSellerIDByRegUserID(regUserID uint) (uint, error) {
	sellerID, err := s.inventoryRepo.GetSellerIDByRegUserID(regUserID)
	if err != nil {
		return 0, errors.New("seller not found")
	}
	return sellerID, nil
}
