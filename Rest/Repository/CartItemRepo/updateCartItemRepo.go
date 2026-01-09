package cartitemrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *cartRepository) UpdateCartItem(item *models.CartItem) error {
	return r.db.Save(item).Error
}
