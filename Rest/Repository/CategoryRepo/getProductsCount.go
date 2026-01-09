package categoryrepo

import models "github.com/saadahmedbd/Treestore/Models"

// GetProductCount counts products in a category (including subcategories)

func (r *categoryRepository) GetProductCount(categoryID uint) (int64, error) {
	// Get all subcategory IDs recursively
	subcategoryIDs, err := r.getSubcategoryIDs(categoryID)
	if err != nil {
		return 0, err
	}

	// Include the category itself
	categoryIDs := append([]uint{categoryID}, subcategoryIDs...)

	// Count products using product repository (if you inject it)
	// For now, keeping the existing direct count:
	var count int64
	err = r.db.Model(&models.Product{}).
		Where("category_id IN ?", categoryIDs).
		Where("is_active = ?", true).
		Count(&count).Error

	return count, err
}
