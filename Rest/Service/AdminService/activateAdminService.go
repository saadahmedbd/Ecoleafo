package adminservice

import "errors"

func (s *adminService) ActivateAdmin(id, requestedBy uint) error {
	// Check if requester is super_admin
	requester, err := s.adminRepo.FindByUserID(requestedBy)
	if err != nil || requester.Role != "super_admin" {
		return errors.New("only super admins can activate admins")
	}

	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return err
	}

	admin.IsActive = true
	return s.adminRepo.Update(admin)
}
