package models

import "time"

// ============================================================================
// SELLER EARNINGS SUMMARY - Cached seller financial data
// ============================================================================
type SellerEarningsSummary struct {
	ID       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID uint `json:"seller_id" gorm:"uniqueIndex;not null"`

	// Sales Metrics
	TotalOrders     int `json:"total_orders" gorm:"default:0"`
	CompletedOrders int `json:"completed_orders" gorm:"default:0"`
	CancelledOrders int `json:"cancelled_orders" gorm:"default:0"`

	// Financial Metrics
	GrossSales      float64 `json:"gross_sales" gorm:"type:decimal(12,2);default:0.00"`      // Total revenue
	TotalCommission float64 `json:"total_commission" gorm:"type:decimal(12,2);default:0.00"` // Platform earnings
	NetEarnings     float64 `json:"net_earnings" gorm:"type:decimal(12,2);default:0.00"`     // Seller earnings

	// Withdrawal Tracking
	TotalWithdrawn   float64 `json:"total_withdrawn" gorm:"type:decimal(12,2);default:0.00"`   // Already paid
	AvailableBalance float64 `json:"available_balance" gorm:"type:decimal(12,2);default:0.00"` // Can withdraw now
	PendingClearance float64 `json:"pending_clearance" gorm:"type:decimal(12,2);default:0.00"` // Not yet available

	LastUpdated time.Time `json:"last_updated"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Seller User `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
}
