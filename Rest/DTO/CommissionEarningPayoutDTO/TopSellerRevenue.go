package commissionearningpayoutdto

type TopSellerByRevenue struct {
	SellerID            uint    `json:"seller_id"`
	SellerName          string  `json:"seller_name"`
	StoreName           string  `json:"store_name"`
	TotalOrders         int     `json:"total_orders"`
	GrossSales          float64 `json:"gross_sales"`
	CommissionGenerated float64 `json:"commission_generated"`
}
