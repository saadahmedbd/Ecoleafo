package adminservice

import (
	"errors"

	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
)

func (s *adminService) UpdateAdminPermissions(id, requestedBy uint, req admin.UpdateAdminPermissionRequest) (*admin.AdminResponse, error) {
	// Check if requester is super_admin
	requester, err := s.adminRepo.FindByUserID(requestedBy)
	if err != nil || requester.Role != "super_admin" {
		return nil, errors.New("only super admins can update permission")
	}
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	//cannot modify super admin permission
	if admin.Role == "super_admin" {
		return nil, errors.New("cannot modify super admin permission")
	}

	if req.CanManageUsers != nil {
		admin.CanManageUsers = *req.CanManageUsers
	}
	if req.CanManageProducts != nil {
		admin.CanManageProducts = *req.CanManageProducts
	}
	if req.CanManageOrders != nil {
		admin.CanManageOrders = *req.CanManageOrders
	}
	if req.CanViewReports != nil {
		admin.CanViewReports = *req.CanViewReports
	}
	if req.CanManageSettings != nil {
		admin.CanManageSettings = *req.CanManageSettings
	}

	if err := s.adminRepo.Update(admin); err != nil {
		return nil, err
	}
	return s.MapToadminsResponse(admin), nil

}
