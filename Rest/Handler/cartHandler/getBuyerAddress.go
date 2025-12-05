package carthandler

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

// getBuyerAddress retrieves the buyer's default address
func getBuyerAddress(buyerID uint) string {
	var address models.Address
	if err := Config.DB.Where("buyer_id = ? AND is_default = ?", buyerID, true).First(&address).Error; err != nil {
		return ""
	}
	// Combine city, district, and state for better zone detection
	result := fmt.Sprintf("%s, %s, %s", address.City, address.District, address.State)
	return result
}
