package adminrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
)

func (r *adminRepository) UpdateLastLogin(userID uint) error {
	now := time.Now()
	return r.db.Model(&models.Admin{}).
		Where("user_id", userID).
		Update("last_login", now).Error
}
