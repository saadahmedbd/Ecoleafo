package selleraccountsetting

// UpdateAccountRequest - Update seller account information (name, phone)
type UpdateAccountRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Phone     string `json:"phone" validate:"required,phone"`
}
