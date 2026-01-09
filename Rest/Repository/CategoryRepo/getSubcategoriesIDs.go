package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// getSubcategoryIDs recursively gets all subcategory IDs
func (r *categoryRepository) getSubcategoryIDs(parentID uint) ([]uint, error) {
	var ids []uint
	var categories []models.Category

	err := r.db.Select("id").
		Where("parent_id = ?", parentID).
		Find(&categories).Error

	if err != nil {
		return nil, err
	}

	for _, cat := range categories {
		ids = append(ids, cat.ID)
		// Recursively get children
		childIDs, err := r.getSubcategoryIDs(cat.ID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, childIDs...)
	}

	return ids, nil
}
