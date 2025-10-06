package guestcartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	address "github.com/saadahmedbd/Treestore/Rest/DTO/Address"
)

func (s *guestcartservice) InitiateCheckout(userID uint, req address.CheckoutInitRequest) error {
	// Check profile completion
	status, err := s.CheckProfileCompletion(userID)
	if err != nil {
		return err
	}

	if !status.IsProfileComplete {
		return fmt.Errorf("profile incomplete: %v", status.MissingFields)
	}

	// Verify addresses exist
	buyer, err := s.buyerRepo.GetBuyerByUserID(userID)
	if err != nil {
		return err
	}

	var shippingAddr, billingAddr models.Address
	if err := Config.DB.Where("id = ? AND buyer_id = ?", req.ShippingAddressID, buyer.ID).First(&shippingAddr).Error; err != nil {
		return fmt.Errorf("invalid shipping address")
	}

	if req.BillingAddressID > 0 {
		if err := Config.DB.Where("id = ? AND buyer_id = ?", req.BillingAddressID, buyer.ID).First(&billingAddr).Error; err != nil {
			return fmt.Errorf("invalid billing address")
		}
	}

	return nil
}
