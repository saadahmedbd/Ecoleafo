package commissionpayoutearningservice

func (s *commissionService) updateSellerEarningsSummary(sellerID uint) error {
	earnings, err := s.commissionRepo.GetSellerEarnings(sellerID)
	if err != nil {
		return err
	}

	// Recalculate totals from order_commissions
	var totals struct {
		TotalOrders      int64
		GrossSales       float64
		TotalCommission  float64
		NetEarnings      float64
		PendingClearance float64
		AvailableBalance float64
	}

	// Get all commissions for this seller
	// You'll need to add a method to get seller's commissions
	// For now, this is pseudo-code

	earnings.TotalOrders = int(totals.TotalOrders)
	earnings.GrossSales = totals.GrossSales
	earnings.TotalCommission = totals.TotalCommission
	earnings.NetEarnings = totals.NetEarnings
	earnings.PendingClearance = totals.PendingClearance
	earnings.AvailableBalance = totals.NetEarnings - earnings.TotalWithdrawn

	return s.commissionRepo.UpdateSellerEarnings(earnings)
}
