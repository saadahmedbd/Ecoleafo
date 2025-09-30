package sellerprofile

type UpdateSellerProfileRequest struct {
	Phone         string `json:"phone"`
	StoreName     string `json:"store_name"`
	StoreDesc     string `json:"store_description"`
	BusinessEmail string `json:"business_email"`
	BusinessType  string `json:"business_type"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	PostalCode    string `json:"postal_code"`
}
