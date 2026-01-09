package commissionearningpayoutdto

type CommissionSettingResponse struct {
	ID          uint    `json:"id"`
	DefaultRate float64 `json:"default_rate"`
	Description string  `json:"description"`
	IsActive    bool    `json:"is_active"`
	UpdatedBy   uint    `json:"updated_by"`
	UpdatedAt   string  `json:"updated_at"`
}
