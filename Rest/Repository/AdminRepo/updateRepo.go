package adminrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *adminRepository) Update(admin *models.Admin) error {
	return r.db.Save(admin).Error

}
