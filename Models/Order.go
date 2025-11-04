package models

import (
	"time"

	"gorm.io/gorm"
)

// // Order represents buyer orders - MVP
type Order struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	OrderNumber   string `json:"order_number" gorm:"size:50;uniqueIndex;not null"`
	BuyerID       uint   `json:"buyer_id" gorm:"not null"`
	Status        string `json:"status" gorm:"size:20;not null;default:'pending'"`         // pending, confirmed, shipped, delivered, cancelled
	PaymentStatus string `json:"payment_status" gorm:"size:20;not null;default:'pending'"` // pending, paid, failed
	PaymentMethod string `json:"payment_method" gorm:"size:50"`                            // cash_on_delivery, bkash, card

	//enhanced financial field
	Subtotal       float64 `json:"subtotal" gorm:"type:decimal(12,2);not null;default:0"`
	ShippingCost   float64 `json:"shipping_cost" gorm:"type:decimal(10,2);default:0"`
	TaxAmount      float64 `json:"tax_amount" gorm:"type:decimal(10,2);default:0"`
	DiscountAmount float64 `json:"discount_amount" gorm:"type:decimal(10,2);default:0"`
	Total          float64 `json:"total" gorm:"type:decimal(12,2);not null"` // Renamed from Total to TotalAmount

	//customer information
	ShippingAddress    string `json:"shipping_address" gorm:"type:text;not null;default:'n/a"`
	BillingAddress     string `json:"billing_address" gorm:"type:text"` // Added billing address
	CustomerEmail      string `json:"customer_email" gorm:"size:100;not null"`
	CustomerPhone      string `json:"customer_phone" gorm:"size:20"`
	CancellationReason string `json:"cancellation_reason" gorm:"type:text"`

	//logistics information
	TrackingNumber string     `json:"tracking_number" gorm:"size:100"` // Added tracking
	ShippedAt      *time.Time `json:"shipped_at"`                      // Added shipping date
	DeliveredAt    *time.Time `json:"delivered_at"`                    // Added delivery date

	Notes string `json:"notes" gorm:"type:text"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Buyer        Buyer          `json:"buyer" gorm:"foreignKey:BuyerID"`
	Reviews      []Review       `json:"reviews" gorm:"foreignKey:OrderID"`
	OrderItems   []OrderItem    `json:"order_items" gorm:"foreignKey:OrderID"`
	OrderHistory []OrderHistory `json:"order_history" gorm:"foreignKey:OrderID"` // Added order history
}
