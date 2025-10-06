package buyeraccount

type CompleteProfileRequest struct {
	Phone string `json:"phone" validate:"required"`
}
