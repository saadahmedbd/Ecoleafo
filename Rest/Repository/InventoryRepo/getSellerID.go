package inventoryrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	// "gorm.io/gorm"
)

func (r *InventoryRepository) GetSellerIDByRegUserID(regUserID uint) (uint, error) {
	var user models.User
	// Try finding by user_id field first (RegUser.ID)
	err := r.db.Where("user_id = ?", regUserID).First(&user).Error
	if err == nil {
		return user.ID, nil
	}

	// Fallback: check if regUserID is actually the User.ID itself
	err = r.db.Where("id = ?", regUserID).First(&user).Error
	if err == nil {
		return user.ID, nil
	}

	return 0, err
}
