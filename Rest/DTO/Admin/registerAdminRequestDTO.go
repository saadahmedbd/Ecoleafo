package admin

type RegisterAdminRequest struct {
	Token           string `json:"token" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	FullName        string `json:"full_name" validate:"required"`
	Phone           string `json:"phone"`
}
