package selleraccount

type CompleteSellerProfileRequest struct {
	//businessInfo
	BusinessEmail string `json:"business_eamil" validate:"email"`
	Phone         string `json:"phone" validate:"required"`
	StoreDesc     string `json:"store_desc"`
	BusinessType  string `json:"business_type" validate:"oneof=individual company nursery"`
	// TaxNumber       string `json:"tax_number"`
	// BusinessLicense string `json:"business_license"`

	//Address info
	Address    string `json:"address" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code" validate:"required"`
}
