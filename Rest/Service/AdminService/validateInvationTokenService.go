package adminservice

import (
	"errors"
	"time"

	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
)

func (s *adminService) ValidateInvitationToken(token string) (*admin.AdminInvitationResponse, error) {
	invitation, err := s.adminRepo.FindInvitationByToken(token)
	if err != nil {
		return nil, errors.New("invalid invitation token")
	}

	if invitation.IsUsed {
		return nil, errors.New("invitation token has already been used")
	}

	if time.Now().After(invitation.ExpiresAt) {
		return nil, errors.New("invitation token has expired")
	}

	return &admin.AdminInvitationResponse{
		ID:         invitation.ID,
		Email:      invitation.Email,
		Token:      invitation.Token,
		Role:       invitation.Role,
		Department: invitation.Department,
		IsUsed:     invitation.IsUsed,
		ExpiresAt:  invitation.ExpiresAt,
		CreatedAt:  invitation.CreatedAt,
		InvitedBy:  "super admin",
	}, nil
}
