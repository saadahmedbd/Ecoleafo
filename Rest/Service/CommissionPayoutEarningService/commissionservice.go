package commissionpayoutearningservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
	commissionpayoutearningrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CommissionPayoutEarningRepo"
	orderrepo "github.com/saadahmedbd/Treestore/Rest/Repository/OrderRepo"
	selleraccountsettingrepo "github.com/saadahmedbd/Treestore/Rest/Repository/sellerAccountSettingRepo"
)

type CommissionService interface {
	// Settings
	GetCommissionSettings() (*commissionearningpayoutdto.CommissionSettingResponse, error)
	UpdateCommissionSettings(req *commissionearningpayoutdto.UpdateCommissionSettingRequest, adminID uint) error

	// Calculate commission when order is created
	CalculateAndCreateCommission(order *models.Order) error

	// Earnings
	GetSellerIDByUserID(userID uint) (uint, error)
	GetSellerEarnings(sellerID uint) (*commissionearningpayoutdto.SellerEarningsResponse, error)
	GetSellerEarningDetails(sellerID uint) (*commissionearningpayoutdto.SellerEarningDetailsResponse, error)
	GetAllSellerEarnings(page, limit int) ([]commissionearningpayoutdto.SellerEarningsResponse, int64, error)

	// Payouts
	RequestPayout(req *commissionearningpayoutdto.CreatePayoutRequest) error
	GetPendingPayouts(page, limit int) ([]commissionearningpayoutdto.PayoutResponse, int64, error)
	ProcessPayout(payoutID uint, req *commissionearningpayoutdto.ProcessPayoutRequest, adminID uint) error
	RejectPayout(payoutID uint, req *commissionearningpayoutdto.RejectPayoutRequest, adminID uint) error
	GetPayoutHistory(sellerID *uint, page, limit int) ([]commissionearningpayoutdto.PayoutResponse, int64, error)

	// Analytics
	GetPlatformEarningsOverview() (*commissionearningpayoutdto.PlatformEarningsOverview, error)
	GetMonthlyRevenue(year int) ([]commissionearningpayoutdto.MonthlyRevenueReport, error)
	GetTopSellers(limit int) ([]commissionearningpayoutdto.TopSellerByRevenue, error)
}

type commissionService struct {
	commissionRepo commissionpayoutearningrepo.CommissionRepository
	orderRepo      orderrepo.OrderRepository
	sellerRepo     selleraccountsettingrepo.SellerAccountSettingRepository
}

func NewCommissionService(
	commissionRepo commissionpayoutearningrepo.CommissionRepository,
	orderRepo orderrepo.OrderRepository,
	sellerRepo selleraccountsettingrepo.SellerAccountSettingRepository,
) CommissionService {
	return &commissionService{
		commissionRepo: commissionRepo,
		orderRepo:      orderRepo,
		sellerRepo:     sellerRepo,
	}
}
