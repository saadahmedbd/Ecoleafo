package reguserrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *regUserRepository) FindByEmail(email string) (*models.RegUser, error) {
	var regUser models.RegUser
	err := r.db.Where("email = ?", email).First(&regUser).Error
	if err != nil {
		return nil, err
	}
	return &regUser, nil
}
