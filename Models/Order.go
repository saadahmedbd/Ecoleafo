package models

import (
	"time"

	"gorm.io/gorm"
)

// // Order represents buyer orders - MVP
type Order struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	OrderNumber string `json:"order_number" gorm:"size:50;uniqueIndex;not null"`
	BuyerID     uint   `json:"buyer_id" gorm:"not null"`
	SellerID    uint   `json:"seller_id" gorm:"null;index"` // Added for multi-seller

	Status        string `json:"status" gorm:"size:20;not null;default:'pending'"`         // pending, confirmed, shipped, delivered, cancelled
	PaymentStatus string `json:"payment_status" gorm:"size:20;not null;default:'pending'"` // pending, paid, failed
	PaymentMethod string `json:"payment_method" gorm:"size:50"`                            // cash_on_delivery, bkash, card

	//enhanced financial field
	Subtotal       float64 `json:"subtotal" gorm:"type:decimal(12,2);not null;default:0"`
	ShippingCost   float64 `json:"shipping_cost" gorm:"type:decimal(10,2);default:0"`
	TaxAmount      float64 `json:"tax_amount" gorm:"type:decimal(10,2);default:0"`
	DiscountAmount float64 `json:"discount_amount" gorm:"type:decimal(10,2);default:0"`
	Total          float64 `json:"total" gorm:"type:decimal(12,2);not null"` // Renamed from Total to TotalAmount

	//commission
	CommissionRate   float64 `json:"commission_rate" gorm:"type:decimal(5,2)"`
	CommissionAmount float64 `json:"commission_amount" gorm:"type:decimal(10,2)"`
	SellerEarnings   float64 `json:"seller_earnings" gorm:"type:decimal(10,2)"`

	//customer information
	ShippingAddress      string `json:"shipping_address" gorm:"type:text;not null;default:'n/a"`
	ShippingPhoneNumber  string `json:"shipping_phone_number" gorm:"size:20"`
	BillingAddress       string `json:"billing_address" gorm:"type:text"` // Added billing address
	BillingPhoneNumber   string `json:"billing_phone_number" gorm:"size:20"`
	CustomerEmail        string `json:"customer_email" gorm:"size:100;not null"`
	CustomerPhone        string `json:"customer_phone" gorm:"size:20"`
	//cancelation/refund
	CancellationReason string  `json:"cancellation_reason" gorm:"type:text"`
	CancelledBy        string  `json:"cancelled_by" gorm:"size:20"` // buyer, seller, admin
	RefundAmount       float64 `json:"refund_amount" gorm:"type:decimal(10,2)"`
	RefundReason       string  `json:"refund_reason" gorm:"type:text"`

	//logistics information
	DeliveryType   string     `json:"delivery_type" gorm:"size:20;default:'home_delivery'"` // home_delivery, pickup_point
	TrackingNumber string     `json:"tracking_number" gorm:"size:100"`                       // Added tracking
	ShippedAt      *time.Time `json:"shipped_at"`                                            // Added shipping date
	DeliveredAt    *time.Time `json:"delivered_at"`                                          // Added delivery date
	OrderDate      time.Time  `json:"order_date"`
	ConfirmedAt    *time.Time `json:"confirmed_at"`

	//gift information
	IsGift      bool   `json:"is_gift" gorm:"default:false"`
	GiftMessage string `json:"gift_message" gorm:"type:text"`
	GiftCharge  float64 `json:"gift_charge" gorm:"type:decimal(10,2);default:0"`

	Notes string `json:"notes" gorm:"type:text"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Buyer        Buyer            `json:"buyer" gorm:"foreignKey:BuyerID"`
	Reviews      []Review         `json:"reviews" gorm:"foreignKey:OrderID"`
	OrderItems   []OrderItem      `json:"order_items" gorm:"foreignKey:OrderID"`
	OrderHistory []OrderHistory   `json:"order_history" gorm:"foreignKey:OrderID"` // Added order
	Commission   *OrderCommission `json:"commission,omitempty" gorm:"foreignKey:OrderID"`
}
