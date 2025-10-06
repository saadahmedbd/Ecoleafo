package address

type CreateAddressRequest struct {
	Type         string `json:"type" validate:"required,oneof=shipping billing"` // shipping, billing
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Company      string `json:"company"`
	AddressLine1 string `json:"address_line_1" validate:"required"`
	AddressLine2 string `json:"address_line_2"`
	Street       string `json:"street"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	District     string `json:"district" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required"`
	Country      string `json:"country"`
	Phone        string `json:"phone" validate:"required"`
	IsDefault    bool   `json:"is_default"`
}
