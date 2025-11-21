package commissionpayoutearningservice

import (
	"time"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
)

func (s *commissionService) GetSellerEarnings(sellerID uint) (*commissionearningpayoutdto.SellerEarningsResponse, error) {
	earnings, err := s.commissionRepo.GetSellerEarnings(sellerID)
	if err != nil {
		return nil, err
	}
	if earnings == nil {
		return nil, nil
	}

	return &commissionearningpayoutdto.SellerEarningsResponse{
		SellerID:         earnings.SellerID,
		SellerName:       earnings.Seller.RegUser.FirstName + " " + earnings.Seller.RegUser.LastName,
		StoreName:        earnings.Seller.StoreName,
		TotalOrders:      earnings.TotalOrders,
		CompletedOrders:  earnings.CompletedOrders,
		GrossSales:       earnings.GrossSales,
		TotalCommission:  earnings.TotalCommission,
		NetEarnings:      earnings.NetEarnings,
		TotalWithdrawn:   earnings.TotalWithdrawn,
		AvailableBalance: earnings.AvailableBalance,
		PendingClearance: earnings.PendingClearance,
		LastUpdated:      earnings.LastUpdated.Format(time.RFC3339),
	}, nil
}

func (s *commissionService) GetAllSellerEarnings(page, limit int) ([]commissionearningpayoutdto.SellerEarningsResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	earningsList, total, err := s.commissionRepo.GetAllSellerEarnings(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var response []commissionearningpayoutdto.SellerEarningsResponse
	for _, earnings := range earningsList {
		response = append(response, commissionearningpayoutdto.SellerEarningsResponse{
			SellerID:         earnings.SellerID,
			SellerName:       earnings.Seller.RegUser.FirstName + " " + earnings.Seller.RegUser.LastName,
			StoreName:        earnings.Seller.StoreName,
			TotalOrders:      earnings.TotalOrders,
			CompletedOrders:  earnings.CompletedOrders,
			GrossSales:       earnings.GrossSales,
			TotalCommission:  earnings.TotalCommission,
			NetEarnings:      earnings.NetEarnings,
			TotalWithdrawn:   earnings.TotalWithdrawn,
			AvailableBalance: earnings.AvailableBalance,
			PendingClearance: earnings.PendingClearance,
			LastUpdated:      earnings.LastUpdated.Format(time.RFC3339),
		})
	}

	return response, total, nil
}
