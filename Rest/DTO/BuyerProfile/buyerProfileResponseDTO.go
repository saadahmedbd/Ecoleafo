package buyerprofile

import (
	"time"

	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
)

type BuyerProfileResponse struct {
	ID               uint                           `json:"id"`
	UserId           uint                           `json:"user_id"`
	Phone            string                         `json:"phone"`
	ProfilePicture   string                         `json:"profile_picture"`
	DefaultAddress   string                         `json:"default_address"`
	IsActive         bool                           `json:"is_active"`
	EmailVerified    bool                           `json:"email_verified"`
	LastOrderAt      *time.Time                     `json:"last_order_at"`
	TotalOrdersCount int                            `json:"total_orders_count"`
	TotalSpent       float64                        `json:"total_spent"`
	CreatedAt        time.Time                      `json:"created_at"`
	UpdatedAt        time.Time                      `json:"updated_at"`
	RegUser          sellerprofile.RegUserBasicInfo `json:"reg_user"`
	DefaultAddr      *AddressInfo                   `json:"default_addr,omitempty"`
	Addresses        []AddressInfo                  `json:"addresses"`
	WishlistCount    int                            `json:"wishlist_count"`
	CartItemCount    int                            `json:"cart_item_count"`
}
