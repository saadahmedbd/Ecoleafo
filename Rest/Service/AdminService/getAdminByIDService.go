package adminservice

import (
	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
)

func (s *adminService) GetAdminByID(id uint) (*admin.AdminResponse, error) {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.MapToadminsResponse(admin), nil
}
