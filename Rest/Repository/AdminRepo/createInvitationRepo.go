package adminrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminRepository) CreateInvitation(invitation *models.AdminInvitation) error {
	return r.db.Create(invitation).Error
}
