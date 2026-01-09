package adminservice

import "errors"

func (s *adminService) DeactivateAdmin(id, requestedBy uint) error {
	//check if requester is super admin
	requester, err := s.adminRepo.FindByUserID(requestedBy)
	if err != nil || requester.Role != "super_admin" {
		return errors.New("only super admin can deactivate admins")
	}
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return err
	}
	//cannot deactivate super admin
	if admin.Role == "super_admin" {
		return errors.New("cannot deactivate super admin")
	}
	admin.IsActive = false
	return s.adminRepo.Update(admin)
}
