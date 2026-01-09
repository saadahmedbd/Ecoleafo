package adminservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
)

func (s *adminService) MapToadminsResponse(admins *models.Admin) *admin.AdminResponse {
	return &admin.AdminResponse{
		ID:                admins.ID,
		UserID:            admins.UserID,
		Email:             admins.RegUser.Email,
		Role:              admins.Role,
		FullName:          admins.FullName,
		Phone:             admins.Phone,
		Department:        admins.Department,
		CanManageUsers:    admins.CanManageUsers,
		CanManageProducts: admins.CanManageProducts,
		CanManageOrders:   admins.CanManageOrders,
		CanViewReports:    admins.CanViewReports,
		CanManageSettings: admins.CanManageSettings,
		IsActive:          admins.IsActive,
		LastLogin:         admins.LastLogin,
		CreatedAt:         admins.CreatedAt,
	}

}
