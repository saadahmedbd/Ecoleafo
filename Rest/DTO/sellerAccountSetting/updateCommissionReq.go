package selleraccountsetting

type UpdateCommissionRequest struct {
	Commission float64 `json:"commission" validate:"required,min=15"`
}
