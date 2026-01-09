package adminrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *adminRepository) GetPendingInvitations() ([]models.AdminInvitation, error) {
	var invitations []models.AdminInvitation
	err := r.db.Preload("InvitedByAdmin.RegUser").
		Where("is_used = false AND expires_at > ?", time.Now()).
		Order("created_at DESC").
		Find(&invitations).Error
	return invitations, err
}
