package commissionearningpayoutdto

type OrderCommissionResponse struct {
	OrderID          uint    `json:"order_id"`
	OrderNumber      string  `json:"order_number"`
	SellerID         uint    `json:"seller_id"`
	SellerName       string  `json:"seller_name"`
	GrossAmount      float64 `json:"gross_amount"`
	CommissionRate   float64 `json:"commission_rate"`
	CommissionAmount float64 `json:"commission_amount"`
	SellerEarnings   float64 `json:"seller_earnings"`
	Status           string  `json:"status"`
	CreatedAt        string  `json:"created_at"`
}
