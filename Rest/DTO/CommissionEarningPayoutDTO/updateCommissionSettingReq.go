package commissionearningpayoutdto

type UpdateCommissionSettingRequest struct {
	DefaultRate float64 `json:"default_rate" binding:"required,min=0,max=100" example:"10.00"`
	Description string  `json:"description" example:"Platform commission rate"`
}
