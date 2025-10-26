package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// Create inserts a new category into the database
func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}
