package commissionpayoutearningservice

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
)

func (s *commissionService) RequestPayout(req *commissionearningpayoutdto.CreatePayoutRequest) error {
	// Validate seller exists and get earnings
	earnings, err := s.commissionRepo.GetSellerEarnings(req.SellerID)
	if err != nil {
		return fmt.Errorf("seller not found: %w", err)
	}

	// Check if seller has sufficient balance
	if earnings.AvailableBalance < req.Amount {
		return errors.New("insufficient balance for withdrawal")
	}

	// Check minimum payout amount (e.g., $50)
	minPayoutAmount := 50.0
	if req.Amount < minPayoutAmount {
		return fmt.Errorf("minimum payout amount is $%.2f", minPayoutAmount)
	}

	// Create payout request
	payout := &models.SellerPayout{
		SellerID:      req.SellerID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		BankName:      req.BankName,
		AccountNumber: req.AccountNumber,
		AccountName:   req.AccountName,
		RequestNote:   req.RequestNote,
		Status:        "pending",
	}

	return s.commissionRepo.CreatePayout(payout)
}

func (s *commissionService) GetPendingPayouts(page, limit int) ([]commissionearningpayoutdto.PayoutResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	payouts, total, err := s.commissionRepo.GetPendingPayouts(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var response []commissionearningpayoutdto.PayoutResponse
	for _, payout := range payouts {
		resp := commissionearningpayoutdto.PayoutResponse{
			ID:            payout.ID,
			SellerID:      payout.SellerID,
			SellerName:    payout.Seller.StoreName,
			Amount:        payout.Amount,
			PaymentMethod: payout.PaymentMethod,
			Status:        payout.Status,
			RequestedAt:   payout.RequestedAt.Format(time.RFC3339),
		}
		response = append(response, resp)
	}

	return response, total, nil
}

func (s *commissionService) ProcessPayout(payoutID uint, req *commissionearningpayoutdto.ProcessPayoutRequest, adminID uint) error {
	// Get payout
	payout, err := s.commissionRepo.GetPayoutByID(payoutID)
	if err != nil {
		return fmt.Errorf("payout not found: %w", err)
	}

	// Check if already processed
	if payout.Status != "pending" {
		return errors.New("payout has already been processed")
	}

	// Update payout with transaction details
	payout.TransactionID = req.TransactionID
	payout.AdminNote = req.AdminNote
	payout.Status = "completed"
	now := time.Now()
	payout.ProcessedAt = &now
	payout.CompletedAt = &now
	payout.ProcessedBy = &adminID

	if err := s.commissionRepo.UpdatePayoutStatus(payoutID, "completed", adminID); err != nil {
		return err
	}

	// Update seller earnings summary
	earnings, err := s.commissionRepo.GetSellerEarnings(payout.SellerID)
	if err != nil {
		return err
	}

	earnings.TotalWithdrawn += payout.Amount
	earnings.AvailableBalance -= payout.Amount

	return s.commissionRepo.UpdateSellerEarnings(earnings)
}

func (s *commissionService) RejectPayout(payoutID uint, req *commissionearningpayoutdto.RejectPayoutRequest, adminID uint) error {
	// Get payout
	payout, err := s.commissionRepo.GetPayoutByID(payoutID)
	if err != nil {
		return fmt.Errorf("payout not found: %w", err)
	}

	// Check if already processed
	if payout.Status != "pending" {
		return errors.New("payout has already been processed")
	}

	// Update payout status
	payout.Status = "rejected"
	payout.RejectionReason = req.RejectionReason
	now := time.Now()
	payout.ProcessedAt = &now
	payout.ProcessedBy = &adminID

	return s.commissionRepo.UpdatePayoutStatus(payoutID, "rejected", adminID)
}

func (s *commissionService) GetPayoutHistory(sellerID *uint, page, limit int) ([]commissionearningpayoutdto.PayoutResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	payouts, total, err := s.commissionRepo.GetPayoutHistory(sellerID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var response []commissionearningpayoutdto.PayoutResponse
	for _, payout := range payouts {
		resp := commissionearningpayoutdto.PayoutResponse{
			ID:              payout.ID,
			SellerID:        payout.SellerID,
			SellerName:      payout.Seller.StoreName,
			Amount:          payout.Amount,
			PaymentMethod:   payout.PaymentMethod,
			Status:          payout.Status,
			TransactionID:   payout.TransactionID,
			RequestedAt:     payout.RequestedAt.Format(time.RFC3339),
			RejectionReason: payout.RejectionReason,
		}

		if payout.ProcessedAt != nil {
			resp.ProcessedAt = payout.ProcessedAt.Format(time.RFC3339)
		}
		if payout.CompletedAt != nil {
			resp.CompletedAt = payout.CompletedAt.Format(time.RFC3339)
		}

		response = append(response, resp)
	}

	return response, total, nil
}
