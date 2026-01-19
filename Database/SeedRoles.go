package database

import (
	"log"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: "buyer", Description: "Buyer role"},
		{Name: "seller", Description: "Seller role"},
		{Name: "admin", Description: "Admin role"},
		{Name: "super_admin", Description: "Super Admin role"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					log.Printf("Failed to create role %s: %v", role.Name, err)
				} else {
					log.Printf("Role %s created successfully", role.Name)
				}
			}
		}
	}
}
