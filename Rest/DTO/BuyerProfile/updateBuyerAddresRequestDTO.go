package buyerprofile

type UpdateBuyerAddressRequest struct {
	FullName     string `json:"full_name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	District     string `json:"district"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	IsDefault    bool   `json:"is_default"`
}
