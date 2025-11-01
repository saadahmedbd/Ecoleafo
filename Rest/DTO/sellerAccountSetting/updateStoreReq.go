package selleraccountsetting

// UpdateStoreRequest - Update store information
type UpdateStoreRequest struct {
	StoreName        string `json:"store_name" validate:"required,min=3,max=100"`
	StoreDescription string `json:"store_description" validate:"required,min=50,max=500"`
	Website          string `json:"website" validate:"omitempty,url"`
	Phone            string `json:"phone" validate:"required,phone"`
	BusinessEmail    string `json:"business_email" validate:"required,email"`
	Address          string `json:"address" validate:"required,min=10,max=255"`
	City             string `json:"city" validate:"required,min=2,max=50"`
	State            string `json:"state" validate:"required,min=2,max=50"`
	Country          string `json:"country" validate:"required,min=2,max=50"`
	PostalCode       string `json:"postal_code" validate:"required,min=3,max=20"`
}
