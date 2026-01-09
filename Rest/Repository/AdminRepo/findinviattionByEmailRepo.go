package adminrepo

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *adminRepository) FindInvitationByEmail(email string) (*models.AdminInvitation, error) {
	var invitataion models.AdminInvitation
	err := r.db.Where("email = ? AND is_used = false AND expires_at > ?", email, time.Now()).
		First(&invitataion).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("email not found")
		}
		return nil, err
	}
	return &invitataion, nil
}
