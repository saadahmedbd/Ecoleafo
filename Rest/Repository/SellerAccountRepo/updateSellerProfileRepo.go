package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) UpdateSellerProfile(seller *models.User) error {
	return r.db.Save(seller).Error
}
