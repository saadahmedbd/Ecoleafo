package buyerprofile

type AddressInfo struct {
	ID           uint   `json:"id"`
	FullName     string `json:"full_name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	District     string `json:"district"`
	Country      string `json:"country"`
	PostalCode   string `json:"postal_code"`
	Phone        string `json:"phone"`
	IsDefault    bool   `json:"is_default"`
}
