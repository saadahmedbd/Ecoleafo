package adminrepo

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *adminRepository) FindByUserID(userID uint) (*models.Admin, error) {
	fmt.Println("Querying admin for user_id:", userID)

	var admin models.Admin
	err := r.db.Preload("RegUser").Where("user_id = ?", userID).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("admin not found")
		}
		return nil, err
	}
	return &admin, nil
}
