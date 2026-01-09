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

	// Calculate real-time earnings from orders
	var totalOrders, completedOrders int64
	var grossSales, totalCommission, netEarnings float64
	
	// Get seller info for commission rate
	seller, err := s.sellerRepo.GetSellerByID(sellerID)
	if err == nil && seller != nil {
		// Count total orders
		s.commissionRepo.GetDB().Table("order_items").
			Select("COUNT(DISTINCT order_id)").
			Where("seller_id = ?", sellerID).
			Scan(&totalOrders)
		
		// Calculate from all order_items for this seller
		type Result struct {
			Count           int64
			GrossSales      float64
			TotalCommission float64
			NetEarnings     float64
		}
		var result Result
		s.commissionRepo.GetDB().Table("order_items").
			Select("COUNT(DISTINCT order_id) as count, COALESCE(SUM(total), 0) as gross_sales, COALESCE(SUM(commission), 0) as total_commission, COALESCE(SUM(seller_earning), 0) as net_earnings").
			Where("seller_id = ?", sellerID).
			Scan(&result)
		
		completedOrders = result.Count
		grossSales = result.GrossSales
		totalCommission = result.TotalCommission
		netEarnings = result.NetEarnings
	}

	return &commissionearningpayoutdto.SellerEarningsResponse{
		SellerID:         earnings.SellerID,
		SellerName:       earnings.Seller.RegUser.FirstName + " " + earnings.Seller.RegUser.LastName,
		StoreName:        earnings.Seller.StoreName,
		CommissionRate:   seller.Commission,
		TotalOrders:      int(totalOrders),
		CompletedOrders:  int(completedOrders),
		GrossSales:       grossSales,
		TotalCommission:  totalCommission,
		NetEarnings:      netEarnings,
		TotalWithdrawn:   earnings.TotalWithdrawn,
		AvailableBalance: netEarnings - earnings.TotalWithdrawn,
		PendingClearance: earnings.PendingClearance,
		LastUpdated:      time.Now().Format(time.RFC3339),
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
		// Calculate real-time data for each seller
		var totalOrders int64
		type Result struct {
			Count           int64
			GrossSales      float64
			TotalCommission float64
			NetEarnings     float64
		}
		var result Result
		
		s.commissionRepo.GetDB().Table("order_items").
			Select("COUNT(DISTINCT order_id) as count, COALESCE(SUM(total), 0) as gross_sales, COALESCE(SUM(commission), 0) as total_commission, COALESCE(SUM(seller_earning), 0) as net_earnings").
			Where("seller_id = ?", earnings.SellerID).
			Scan(&result)
		
		totalOrders = result.Count
		
		response = append(response, commissionearningpayoutdto.SellerEarningsResponse{
			SellerID:         earnings.SellerID,
			SellerName:       earnings.Seller.RegUser.FirstName + " " + earnings.Seller.RegUser.LastName,
			StoreName:        earnings.Seller.StoreName,
			CommissionRate:   earnings.Seller.Commission,
			TotalOrders:      int(totalOrders),
			CompletedOrders:  int(result.Count),
			GrossSales:       result.GrossSales,
			TotalCommission:  result.TotalCommission,
			NetEarnings:      result.NetEarnings,
			TotalWithdrawn:   earnings.TotalWithdrawn,
			AvailableBalance: result.NetEarnings - earnings.TotalWithdrawn,
			PendingClearance: earnings.PendingClearance,
			LastUpdated:      time.Now().Format(time.RFC3339),
		})
	}

	return response, total, nil
}
