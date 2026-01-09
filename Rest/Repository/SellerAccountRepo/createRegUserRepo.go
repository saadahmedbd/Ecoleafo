package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) CreateRegUser(regUser *models.RegUser) error {
	return r.db.Create(regUser).Error
}
