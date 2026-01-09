package adminservice

import admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"

func (s *adminService) GetAllAdmins(page, limit int) ([]admin.AdminResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	admins, total, err := s.adminRepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]admin.AdminResponse, len(admins))
	for i, admin := range admins {
		responses[i] = *s.MapToadminsResponse(&admin)
	}

	return responses, total, nil
}
