package models

import "time"

// ============================================================================
// PLATFORM REVENUE SUMMARY - Overall platform earnings tracking
// ============================================================================
type PlatformRevenueSummary struct {
	ID    uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Year  int  `json:"year" gorm:"not null"`
	Month int  `json:"month" gorm:"not null"`

	// Order Metrics
	TotalOrders     int `json:"total_orders" gorm:"default:0"`
	CompletedOrders int `json:"completed_orders" gorm:"default:0"`

	// Revenue Metrics
	GrossRevenue        float64 `json:"gross_revenue" gorm:"type:decimal(15,2);default:0.00"`    // Total sales
	TotalCommission     float64 `json:"total_commission" gorm:"type:decimal(15,2);default:0.00"` // Platform earnings
	TotalSellerEarnings float64 `json:"total_seller_earnings" gorm:"type:decimal(15,2);default:0.00"`

	// Payout Metrics
	TotalPayouts   float64 `json:"total_payouts" gorm:"type:decimal(15,2);default:0.00"`
	PendingPayouts float64 `json:"pending_payouts" gorm:"type:decimal(15,2);default:0.00"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Add unique index for year+month
func (PlatformRevenueSummary) TableName() string {
	return "platform_revenue_summaries"
}
