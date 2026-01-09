package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Delete soft deletes a category (sets deleted_at timestamp)
func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
