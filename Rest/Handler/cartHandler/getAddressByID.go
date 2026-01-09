package carthandler

import (
	"fmt"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

// getAddressByID retrieves address by ID and returns formatted string
func getAddressByID(addressIDStr string) string {
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		return ""
	}

	var address models.Address
	if err := Config.DB.Where("id = ?", uint(addressID)).First(&address).Error; err != nil {
		return ""
	}

	return fmt.Sprintf("%s, %s, %s", address.City, address.District, address.State)
}
