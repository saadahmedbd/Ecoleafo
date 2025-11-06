package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents sellers who can manage products
type User struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID uint `json:"role_id" gorm:"not null"`
	UserId uint `json:"user_id" gorm:"uniqueindex"` // References RegUser.ID

	//business info
	BusinessEmail string `json:"business_email" gorm:"uniqueIndex;size:100;default:'N/A'"`
	Password      string `json:"-" gorm:"size:255;not null"` // Hidden in JSON
	Phone         string `json:"phone" gorm:"size:20"`
	StoreName     string `json:"store_name" gorm:"size:100;not null"`    // Seller's store name
	StoreSlug     string `json:"store_slug" gorm:"size:100;uniqueIndex"` // Added for SEO
	StoreDesc     string `json:"store_description" gorm:"type:text"`     // Store description
	StoreLogo     string `json:"store_logo" gorm:"size:500"`             // Added store logo
	StoreBanner   string `json:"store_banner" gorm:"size:500"`           // Added store banner
	Website       string `json:"website" gorm:"size:100;default:'N/A'"`
	//Business detils
	BusinessType    string `json:"business_type" gorm:"size:50;default:'individual'"` // individual, company, nursery
	TaxNumber       string `json:"tax_number" gorm:"size:50"`                         // Added tax info
	BusinessLicense string `json:"business_license" gorm:"size:100"`                  // Added license
	// Financial Information

	TotalSales    float64 `json:"total_sales" gorm:"type:decimal(12,2);default:0"`    // Increased precision
	TotalEarnings float64 `json:"total_earnings" gorm:"type:decimal(12,2);default:0"` // Added earnings
	TotalOrders   int     `json:"total_orders" gorm:"default:0"`
	AverageRating float64 `json:"average_rating" gorm:"type:decimal(3,2);default:0"` // Added rating
	Commission    float64 `json:"commission" gorm:"type:decimal(5,2);default:10" `

	//address and location
	Address    string `json:"address" gorm:"type:text;default:'N/A';not null"` // Seller's address
	City       string `json:"city" gorm:"size:50"`                             // city
	State      string `json:"state" gorm:"size:50"`                            //  state
	Country    string `json:"country" gorm:"size:50;default:'Bangladesh'"`     //  country
	PostalCode string `json:"postal_code" gorm:"size:20"`                      // postal code

	// Status & Verification
	Status         string `json:"status" gorm:"size:20;default:'pending'"` // pending, approved, rejected, suspended
	ApprovalStatus string `json:"approval_status" gorm:"size:20;default:'pending'"`

	IsActive          bool       `json:"is_active" gorm:"default:true"`
	IsVerified        bool       `json:"is_verified" gorm:"default:false"`  // Seller verification
	IsApproved        bool       `json:"is_approved" gorm:"default:false"`  //  approval status
	ApprovedAt        *time.Time `json:"approved_at"`                       //  approval date
	RejectedAt        *time.Time `json:"rejected_at"`                       //  rejection date
	RejectionReason   string     `json:"rejection_reason" gorm:"type:text"` // rejection reason
	IsProfileComplete bool       `json:"is_profile_complete" gorm:"default:false"`
	//after compelete profile
	HasBusinessInfo  bool    `json:"has_business_info" gorm:"default:false"`
	HasAddress       bool    `json:"has_address" gorm:"default:false"`
	HasPaymentMethod bool    `json:"has_payment_method" gorm:"default:false"`
	CanAddProducts   bool    `json:"can_add_products" gorm:"default:false"`
	MissingFields    *string `json:"missing_fields" gorm:"type:text"`
	NextStep         string  `json:"next_step" gorm:"size:50;default:'wait_approval'"`

	ApprovedBy *uint `json:"approved_by"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Role             Role                  `json:"role" gorm:"foreignKey:RoleID"`
	Products         []Product             `json:"products" gorm:"foreignKey:SellerID;references:ID"` // Seller's products
	RegUser          *RegUser              `json:"reg_user" gorm:"foreignKey:UserId;references:ID"`
	PaymentMethods   []SellerPaymentMethod `json:"payment_methods" gorm:"foreignKey:SellerID"`   //
	SellerCategories []SellerCategory      `json:"seller_categories" gorm:"foreignKey:SellerID"` //
	ApprovedByAdmin  *Admin                `json:"approved_by_admin,omitempty" gorm:"foreignKey:ApprovedBy;references:ID"`
}
