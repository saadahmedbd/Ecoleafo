package orderhandler

import (
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func getBuyerID(w http.ResponseWriter, r *http.Request) (uint, error) {
	userIDStr := r.Header.Get("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return 0, err
	}

	var buyer models.Buyer
	if err := Config.DB.Where("user_id = ?", uint(userID)).First(&buyer).Error; err != nil {
		return 0, err
	}

	return buyer.ID, nil
}
