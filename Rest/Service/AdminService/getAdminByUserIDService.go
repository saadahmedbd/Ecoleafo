package adminservice

import (
	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
)

func (s *adminService) GetAdminByUserID(userID uint) (*admin.AdminResponse, error) {
	admin, err := s.adminRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return s.MapToadminsResponse(admin), nil
}
