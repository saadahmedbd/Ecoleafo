package adminrepo

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *adminRepository) FindInvitationByToken(token string) (*models.AdminInvitation, error) {
	var invitation models.AdminInvitation

	err := r.db.Where("token = ? ", token).First(&invitation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("inviattaion not found")
		}
		return nil, err
	}
	return &invitation, nil
}
