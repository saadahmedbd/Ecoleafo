package adminrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminRepository) DeleteInvitation(id uint) error {
	return r.db.Delete(&models.AdminInvitation{}, id).Error
}
