package commissionpayoutearningservice

import models "github.com/saadahmedbd/Treestore/Models"

// ============================================================================
// CALCULATE COMMISSION (Called when order is created/completed)
// ============================================================================
func (s *commissionService) CalculateAndCreateCommission(order *models.Order) error {
	// Get commission rate (from seller or default)
	settings, err := s.commissionRepo.GetCommissionSettings()
	if err != nil {
		return err
	}

	// Get seller to check custom commission rate
	seller, err := s.sellerRepo.GetSellerByID(order.SellerID)
	if err != nil {
		return err
	}

	commissionRate := seller.Commission
	if commissionRate == 0 {
		commissionRate = settings.DefaultRate
	}

	// Calculate amounts
	grossAmount := order.Total
	commissionAmount := grossAmount * (commissionRate / 100.0)
	sellerEarnings := grossAmount - commissionAmount

	// Create commission record
	commission := &models.OrderCommission{
		OrderID:          order.ID,
		SellerID:         order.SellerID,
		GrossAmount:      grossAmount,
		CommissionRate:   commissionRate,
		CommissionAmount: commissionAmount,
		SellerEarnings:   sellerEarnings,
		Status:           "pending",
	}

	if err := s.commissionRepo.CreateOrderCommission(commission); err != nil {
		return err
	}

	// Update seller earnings summary
	return s.updateSellerEarningsSummary(order.SellerID)
}
