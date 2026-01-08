package adminservice

import (
	"errors"
	"time"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *adminService) RegisterAdmin(req admin.RegisterAdminRequest) (*admin.RegisterAdminResponse, error) {
	//validate invitation token
	invitation, err := s.adminRepo.FindInvitationByToken(req.Token)
	if err != nil {
		return nil, errors.New("invalid inviatation token")
	}
	if invitation.IsUsed {
		return nil, errors.New("invitation token already used")
	}
	if time.Now().After(invitation.ExpiresAt) {
		return nil, errors.New("invitation token expired")
	}
	//check email already exists
	existingUser, _ := s.reguserrepo.FindByEmail(invitation.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}
	//hashpassword
	hashpassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	//create user account
	user := &models.RegUser{
		Email:    invitation.Email,
		Password: hashpassword,
		Role:     "admin",
		IsActive: true,
	}
	if err := Config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	//create admin profile
	admins := &models.Admin{
		UserID:            user.ID,
		Role:              invitation.Role,
		FullName:          req.FullName,
		Phone:             req.Phone,
		Department:        invitation.Department,
		CanManageUsers:    true,
		CanManageProducts: true,
		CanManageOrders:   true,
		CanViewReports:    true,
		CanManageSettings: invitation.Role == "super_admin",
		IsActive:          true,
	}
	if err := s.adminRepo.Create(admins); err != nil {
		Config.DB.Delete(user)
		return nil, err
	}
	//mark inviateion as used
	invitation.IsUsed = true
	s.adminRepo.UpdateInvitation(invitation)

	//generate jwt token
	token, err := util.CreateJwt(user.ID, admins.FullName, "", []string{"admin"}, admins.ID, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &admin.RegisterAdminResponse{
		Admin: &admin.AdminResponse{
			ID:                admins.ID,
			UserID:            admins.UserID,
			Email:             user.Email,
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
			CreatedAt:         admins.CreatedAt,
		},
		Token: token,
	}, nil
}
