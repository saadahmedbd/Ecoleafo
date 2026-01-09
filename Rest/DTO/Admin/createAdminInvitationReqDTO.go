package admin

type CreateAdminInvitationRequest struct {
	Email         string `json:"email" binding:"required,email"`
	Role          string `json:"role" validate:"required,oneof=super_admin admin"`
	Department    string `json:"department"`
	ExpireInHours int    `json:"expire_in_hours" validate:"required,min=1,max=168"` //max 7 days
}
