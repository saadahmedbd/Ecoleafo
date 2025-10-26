package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Update updates an existing category
func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}
