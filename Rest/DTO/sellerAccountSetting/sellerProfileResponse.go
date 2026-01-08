package selleraccountsetting

import "time"

// SellerProfileResponse - Complete seller profile
type SellerProfileResponse struct {
	// Basic Info
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	ProfilePhoto string `json:"profile_photo"`

	// Store Info
	StoreName        string `json:"store_name"`
	StoreSlug        string `json:"store_slug"`
	StoreDescription string `json:"store_description"`
	StoreLogo        string `json:"store_logo"`
	StoreBanner      string `json:"store_banner"`
	Website          string `json:"website,omitempty"`

	// Business Info
	BusinessEmail   string `json:"business_email"`
	BusinessType    string `json:"business_type"`
	TaxNumber       string `json:"tax_number,omitempty"`
	BusinessLicense string `json:"business_license,omitempty"`

	// Address
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`

	// Statistics
	TotalSales    float64 `json:"total_sales"`
	TotalEarnings float64 `json:"total_earnings"`
	TotalOrders   int     `json:"total_orders"`
	AverageRating float64 `json:"average_rating"`
	Commission    float64 `json:"commission"`

	// Status
	IsActive        bool       `json:"is_active"`
	IsVerified      bool       `json:"is_verified"`
	IsApproved      bool       `json:"is_approved"`
	ApprovedAt      *time.Time `json:"approved_at,omitempty"`
	RejectedAt      *time.Time `json:"rejected_at,omitempty"`
	RejectionReason string     `json:"rejection_reason,omitempty"`

	// Security
	TwoFactorEnabled bool `json:"two_factor_enabled"`

	// Policies (if stored in separate table, otherwise add to User model)
	ReturnPolicy   string `json:"return_policy,omitempty"`
	ShippingPolicy string `json:"shipping_policy,omitempty"`
	FAQ            string `json:"faq,omitempty"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
