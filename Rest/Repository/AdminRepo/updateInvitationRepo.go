package adminrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminRepository) UpdateInvitation(invitation *models.AdminInvitation) error {
	return r.db.Save(invitation).Error
}
