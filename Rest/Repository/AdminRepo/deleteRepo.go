package adminrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminRepository) Delete(id uint) error {
	return r.db.Delete(&models.Admin{}, id).Error
}
