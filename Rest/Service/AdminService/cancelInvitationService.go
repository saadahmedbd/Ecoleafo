package adminservice

import "errors"

func (s *adminService) CancelInvitation(invitationID, requestedBy uint) error {
	// Check if requester is super_admin
	requester, err := s.adminRepo.FindByUserID(requestedBy)
	if err != nil || requester.Role != "super_admin" {
		return errors.New("only super admins can cancel invitations")
	}
	return s.adminRepo.DeleteInvitation(invitationID)

}
