package adminservice

import (
	"errors"

	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	admin "github.com/saadahmedbd/Treestore/Rest/DTO/Admin"
)

func (s *adminService) CreateInvitation(adminID uint, req admin.CreateAdminInvitationRequest) (*admin.AdminInvitationResponse, error) {
	//check admin exist and is super admin
	admins, err := s.adminRepo.FindByUserID(adminID)
	if err != nil {
		return nil, errors.New("unauthorized admin not founc")
	}
	if admins.Role != "super_admin" {
		return nil, errors.New("only super admin can invite new admins")
	}
	//check if email already has a pending invitation
	existingInvitation, _ := s.adminRepo.FindInvitationByEmail(req.Email)
	if existingInvitation != nil {
		return nil, errors.New("an active invitation already exist for this email")
	}
	//check if email already has a user
	existingUser, _ := s.reguserrepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("a user already exist for this email")
	}
	//generate secure token
	token, err := generateSecureToken(32)
	if err != nil {
		return nil, errors.New("error generating token")
	}
	//set expiration time
	expiresIn := req.ExpireInHours
	if expiresIn == 0 {
		expiresIn = 72 //default 3 days
	}
	invitation := &models.AdminInvitation{
		Email:      req.Email,
		Token:      token,
		Role:       req.Role,
		Department: req.Department,
		InvitedBy:  admins.ID,
		IsUsed:     false,
		ExpiresAt:  time.Now().Add(time.Duration(expiresIn) * time.Hour),
	}
	if err := s.adminRepo.CreateInvitation(invitation); err != nil {
		return nil, err
	}

	return &admin.AdminInvitationResponse{
		ID:         invitation.ID,
		Email:      invitation.Email,
		Token:      invitation.Token,
		Role:       invitation.Role,
		Department: invitation.Department,
		IsUsed:     invitation.IsUsed,
		ExpiresAt:  invitation.ExpiresAt,
		CreatedAt:  invitation.CreatedAt,
		InvitedBy:  admins.FullName,
	}, nil

}
