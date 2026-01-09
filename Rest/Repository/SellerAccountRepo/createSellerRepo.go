package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) CreateSeller(seller *models.User) error {
	return r.db.Create(seller).Error
}
