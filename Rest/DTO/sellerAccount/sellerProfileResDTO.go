package selleraccount

import (
	"time"

	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
)

type SellerProfileResponse struct {
	ID              uint                           `json:"id"`
	UserId          uint                           `json:"user_id"`
	BusinessEmail   string                         `json:"business_email"`
	Phone           string                         `json:"phone"`
	StoreName       string                         `json:"store_name"`
	StoreSlug       string                         `json:"store_slug"`
	StoreDesc       string                         `json:"store_description"`
	StoreLogo       string                         `json:"store_logo"`
	StoreBanner     string                         `json:"store_banner"`
	BusinessType    string                         `json:"business_type"`
	TaxNumber       string                         `json:"tax_number"`
	BusinessLicense string                         `json:"business_license"`
	TotalSales      float64                        `json:"total_sales"`
	TotalEarnings   float64                        `json:"total_earnings"`
	TotalOrders     int                            `json:"total_orders"`
	AverageRating   float64                        `json:"average_rating"`
	Commission      float64                        `json:"commission"`
	Address         string                         `json:"address"`
	City            string                         `json:"city"`
	State           string                         `json:"state"`
	Country         string                         `json:"country"`
	PostalCode      string                         `json:"postal_code"`
	IsActive        bool                           `json:"is_active"`
	IsVerified      bool                           `json:"is_verified"`
	IsApproved      bool                           `json:"is_approved"`
	ApprovedAt      *time.Time                     `json:"approved_at"`
	RejectedAt      *time.Time                     `json:"rejected_at"`
	RejectionReason string                         `json:"rejection_reason"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
	ProductCount    int                            `json:"product_count"`
	PaymentMethods  []PaymentMethodInfo            `json:"payment_methods"`
	RegUser         sellerprofile.RegUserBasicInfo `json:"reg_user"`
}
