package models

import "time"

// ============================================================================
// ORDER COMMISSION - Track commission for each order
// ============================================================================
type OrderCommission struct {
	ID       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID  uint `json:"order_id" gorm:"not null;index"`
	SellerID uint `json:"seller_id" gorm:"not null;index"`

	// Financial Breakdown
	GrossAmount      float64 `json:"gross_amount" gorm:"type:decimal(10,2);not null"`      // Order total
	CommissionRate   float64 `json:"commission_rate" gorm:"type:decimal(5,2);not null"`    // % taken
	CommissionAmount float64 `json:"commission_amount" gorm:"type:decimal(10,2);not null"` // Platform earnings
	SellerEarnings   float64 `json:"seller_earnings" gorm:"type:decimal(10,2);not null"`   // What seller gets

	// Status Tracking
	Status    string     `json:"status" gorm:"size:20;default:'pending'"` // pending, cleared, paid_out
	ClearedAt *time.Time `json:"cleared_at"`                              // When funds are available for withdrawal
	PaidOutAt *time.Time `json:"paid_out_at"`                             // When paid to seller

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Order  Order `json:"order" gorm:"foreignKey:OrderID;references:ID"`
	Seller User  `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
}
