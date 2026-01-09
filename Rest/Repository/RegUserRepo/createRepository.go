package reguserrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *regUserRepository) Create(user *models.RegUser) error {
	return r.db.Save(user).Error
}
