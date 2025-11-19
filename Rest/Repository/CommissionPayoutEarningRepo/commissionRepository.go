package commissionpayoutearningrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type CommissionRepository interface {
	// Commission Settings
	GetCommissionSettings() (*models.CommissionSetting, error)
	UpdateCommissionSettings(settings *models.CommissionSetting) error

	// Order Commission
	CreateOrderCommission(commission *models.OrderCommission) error
	GetOrderCommission(orderID uint) (*models.OrderCommission, error)
	UpdateCommissionStatus(orderID uint, status string) error

	// Seller Earnings
	GetSellerEarnings(sellerID uint) (*models.SellerEarningsSummary, error)
	UpdateSellerEarnings(earnings *models.SellerEarningsSummary) error
	GetAllSellerEarnings(page, limit int) ([]models.SellerEarningsSummary, int64, error)

	// Payouts
	CreatePayout(payout *models.SellerPayout) error
	GetPayoutByID(id uint) (*models.SellerPayout, error)
	GetPendingPayouts(page, limit int) ([]models.SellerPayout, int64, error)
	GetPayoutHistory(sellerID *uint, page, limit int) ([]models.SellerPayout, int64, error)
	UpdatePayoutStatus(id uint, status string, processedBy uint) error

	// Platform Analytics
	GetPlatformEarningsOverview() (map[string]interface{}, error)
	GetMonthlyRevenue(year int) ([]map[string]interface{}, error)
	GetTopSellersByRevenue(limit int) ([]map[string]interface{}, error)
}

type commissionRepository struct {
	db *gorm.DB
}

func NewCommissionRepository(db *gorm.DB) CommissionRepository {
	return &commissionRepository{db: db}
}
