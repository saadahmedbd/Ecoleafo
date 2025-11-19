package orderhandler

import (
	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func getAdminID(regUserID uint) (uint, error) {

	var buyer models.Admin
	if err := Config.DB.Where("user_id = ?", regUserID).First(&buyer).Error; err != nil {
		return 0, err
	}
	return buyer.ID, nil
}
