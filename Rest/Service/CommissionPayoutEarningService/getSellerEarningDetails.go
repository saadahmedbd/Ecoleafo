package commissionpayoutearningservice

import (
	"fmt"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
)

func (s *commissionService) GetSellerEarningDetails(sellerID uint) (*commissionearningpayoutdto.SellerEarningDetailsResponse, error) {
	earnings, orders, err := s.commissionRepo.GetSellerEarningDetails(sellerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get seller earnings: %v", err)
	}

	seller := earnings.Seller
	var sellerName string
	var email string
	if seller.RegUser != nil {
		sellerName = seller.RegUser.FirstName + " " + seller.RegUser.LastName
		email = seller.RegUser.Email
	}

	recentOrders := make([]commissionearningpayoutdto.RecentOrderSummary, 0)
	for _, order := range orders {
		commission := order.Total * (seller.Commission / 100)
		sellerEarned := order.Total - commission

		recentOrders = append(recentOrders, commissionearningpayoutdto.RecentOrderSummary{
			OrderID:      order.ID,
			OrderNumber:  order.OrderNumber,
			OrderDate:    order.CreatedAt.Format("2006-01-02 15:04:05"),
			OrderTotal:   order.Total,
			Commission:   commission,
			SellerEarned: sellerEarned,
			Status:       order.Status,
		})
	}

	return &commissionearningpayoutdto.SellerEarningDetailsResponse{
		SellerID:         earnings.SellerID,
		SellerName:       sellerName,
		StoreName:        seller.StoreName,
		Email:            email,
		Phone:            seller.Phone,
		CommissionRate:   seller.Commission,
		TotalOrders:      earnings.TotalOrders,
		CompletedOrders:  earnings.CompletedOrders,
		CancelledOrders:  earnings.CancelledOrders,
		GrossSales:       earnings.GrossSales,
		TotalCommission:  earnings.TotalCommission,
		NetEarnings:      earnings.NetEarnings,
		TotalWithdrawn:   earnings.TotalWithdrawn,
		AvailableBalance: earnings.AvailableBalance,
		PendingClearance: earnings.PendingClearance,
		LastUpdated:      earnings.LastUpdated.Format("2006-01-02 15:04:05"),
		RecentOrders:     recentOrders,
	}, nil
}
