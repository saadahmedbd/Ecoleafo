package admin

import "time"

type AdminInvitationResponse struct {
	ID         uint      `json:"id"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
	Role       string    `json:"role"`
	Department string    `json:"department"`
	IsUsed     bool      `json:"is_used"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
	InvitedBy  string    `json:"invited_by"`
}
