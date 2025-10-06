package guestcartservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	buyeraccount "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerAccount"
)

func (s *guestcartservice) CompleteProfile(userID uint, req buyeraccount.CompleteProfileRequest) error {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userID)
	if err != nil {
		// If buyer doesn't exist, create one
		if err.Error() == "buyer not found" {
			buyer = &models.Buyer{
				UserId: userID,
				Phone:  req.Phone,
				RoleID: 2, // Assuming 2 is buyer role ID
			}
			return s.buyerRepo.UpdateBuyerProfile(buyer)
		}
		return err
	}
	buyer.Phone = req.Phone
	return s.profileRepo.UpdateBuyerProfile(buyer)
}
