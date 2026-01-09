package adminrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *adminRepository) FindAll(page, limit int) ([]models.Admin, int64, error) {
	var admins []models.Admin
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.Admin{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("RegUser").
		Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&admins).Error

	return admins, total, err
}
