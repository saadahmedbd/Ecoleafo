package adminservice

import (
	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
	adminrepo "github.com/saadahmedbd/Treestore/Rest/Repository/AdminRepo"
	reguserrepo "github.com/saadahmedbd/Treestore/Rest/Repository/RegUserRepo"
)

type Adminservice interface {
	CreateInvitation(adminID uint, req admin.CreateAdminInvitationRequest) (*admin.AdminInvitationResponse, error)
	RegisterAdmin(req admin.RegisterAdminRequest) (*admin.RegisterAdminResponse, error)
	GetAdminByID(id uint) (*admin.AdminResponse, error)
	GetAdminByUserID(userID uint) (*admin.AdminResponse, error)
	GetAllAdmins(page, limit int) ([]admin.AdminResponse, int64, error)
	UpdateAdmin(id uint, req admin.UpdateAdminRequest) (*admin.AdminResponse, error)
	UpdateAdminPermissions(id, requestedBy uint, req admin.UpdateAdminPermissionRequest) (*admin.AdminResponse, error)
	DeactivateAdmin(id, requestedBy uint) error
	ActivateAdmin(id, requestedBy uint) error
	GetPendingInvitations() ([]admin.AdminInvitationResponse, error)
	CancelInvitation(invitationID, requestedBy uint) error
	ValidateInvitationToken(token string) (*admin.AdminInvitationResponse, error)
}
type adminService struct {
	adminRepo   adminrepo.AdminRepo
	reguserrepo reguserrepo.RegUserRepository
}

func NewAdminService(adminRepo adminrepo.AdminRepo, reguserrepo reguserrepo.RegUserRepository) Adminservice {
	return &adminService{
		adminRepo:   adminRepo,
		reguserrepo: reguserrepo,
	}
}
