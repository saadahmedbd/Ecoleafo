package adminservice

import admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"

func (s *adminService) UpdateAdmin(id uint, req admin.UpdateAdminRequest) (*admin.AdminResponse, error) {
	admins, err := s.adminRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if req.FullName != "" {
		admins.FullName = req.FullName
	}
	if req.Phone != "" {
		admins.Phone = req.Phone
	}
	if req.Department != "" {
		admins.Department = req.Department
	}
	if err := s.adminRepo.Update(admins); err != nil {
		return nil, err
	}
	return s.MapToadminsResponse(admins), nil
}
