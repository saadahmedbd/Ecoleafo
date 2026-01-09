package adminservice

import admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"

func (s *adminService) GetPendingInvitations() ([]admin.AdminInvitationResponse, error) {
	invitations, err := s.adminRepo.GetPendingInvitations()
	if err != nil {
		return nil, err
	}
	response := make([]admin.AdminInvitationResponse, len(invitations))
	for i, inv := range invitations {
		response[i] = admin.AdminInvitationResponse{
			ID:         inv.ID,
			Email:      inv.Email,
			Token:      inv.Token,
			Role:       inv.Role,
			Department: inv.Department,
			IsUsed:     inv.IsUsed,
			ExpiresAt:  inv.ExpiresAt,
			CreatedAt:  inv.CreatedAt,
			InvitedBy:  inv.InvitedByAdmin.RegUser.FirstName,
		}
	}
	return response, nil
}
