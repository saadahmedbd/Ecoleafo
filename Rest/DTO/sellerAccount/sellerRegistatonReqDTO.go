package selleraccount

type SellerRegistationRequest struct {
	Email           string  `json:"email" validate:"required,email"`
	Password        string  `json:"password" validate:"required,min=8"`
	ConfirmPassword string  `json:"confirm_password" validate:"required"`
	FirstName       string  `json:"first_name" validate:"required"`
	LastName        string  `json:"last_name" validate:"required"`
	StoreName       string  `json:"store_name" validate:"required"`
	Phone           string  `json:"phone" validate:"required"`
	Commission      float64 `json:"commission"`
	AgreeToTerms    bool    `json:"agree_to_terms" validate:"required"`
}
