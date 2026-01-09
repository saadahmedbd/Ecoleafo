package adminrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminRepository) Create(admin *models.Admin) error {
	return r.db.Create(admin).Error
}
