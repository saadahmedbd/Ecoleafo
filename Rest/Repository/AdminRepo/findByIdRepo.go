package adminrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *adminRepository) FindByID(id uint) (*models.Admin, error) {

	var admin models.Admin

	err := r.db.Preload("RegUser").First(&admin, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin not found")
		}
		return nil, err

	}
	return &admin, nil
}
