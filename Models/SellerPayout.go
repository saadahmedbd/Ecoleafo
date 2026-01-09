package models

import (
	"time"

	"gorm.io/gorm"
)

// ============================================================================
// SELLER PAYOUT - Withdrawal/payout requests
// ============================================================================
type SellerPayout struct {
	ID       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID uint `json:"seller_id" gorm:"not null;index"`

	// Payout Details
	Amount   float64 `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency string  `json:"currency" gorm:"size:3;default:'USD'"`

	// Payment Information
	PaymentMethod string `json:"payment_method" gorm:"size:50"` // bank_transfer, paypal, stripe
	BankName      string `json:"bank_name" gorm:"size:100"`
	AccountNumber string `json:"account_number" gorm:"size:100"`
	AccountName   string `json:"account_name" gorm:"size:100"`
	RoutingNumber string `json:"routing_number" gorm:"size:50"`
	PaypalEmail   string `json:"paypal_email" gorm:"size:100"`

	// Transaction Details
	TransactionID  string  `json:"transaction_id" gorm:"size:100"` // Bank/payment gateway ref
	TransactionFee float64 `json:"transaction_fee" gorm:"type:decimal(10,2);default:0.00"`
	NetAmount      float64 `json:"net_amount" gorm:"type:decimal(10,2)"` // After fees

	// Status & Approval
	Status          string `json:"status" gorm:"size:20;default:'pending'"` // pending, processing, completed, failed, rejected
	RequestNote     string `json:"request_note" gorm:"type:text"`           // Seller's note
	AdminNote       string `json:"admin_note" gorm:"type:text"`             // Admin's note
	RejectionReason string `json:"rejection_reason" gorm:"type:text"`

	// Timestamps
	RequestedAt time.Time  `json:"requested_at"`
	ProcessedAt *time.Time `json:"processed_at"`
	CompletedAt *time.Time `json:"completed_at"`

	// Admin Tracking
	ProcessedBy *uint `json:"processed_by"` // Admin who processed

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Seller           User   `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
	ProcessedByAdmin *Admin `json:"processed_by_admin,omitempty" gorm:"foreignKey:ProcessedBy;references:ID"`
}
