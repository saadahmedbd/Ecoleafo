package database

import (
	"fmt"
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func MigrateRegUser(db *gorm.DB) error {
	err := db.AutoMigrate(&models.RegUser{})
	if err != nil {
		return fmt.Errorf("failed to migrate RegUser: %w", err)
	}
	fmt.Println("RegUser model migrated successfully!")
	return nil
}