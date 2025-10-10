package cartitem

type ApplyCouponRequest struct {
	CouponCode string `json:"coupon_code" validate:"required"`
}
