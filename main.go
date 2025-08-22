package main

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func main() {
	Config.Connect()
	// Run Auto Migration
	err := Config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderItem{},
		&models.Buyer{},
		&models.CartItem{},
		&models.Review{},
		&models.Role{},
	)
	if err != nil {
		fmt.Println("❌ Migration failed:", err)
	} else {
		fmt.Println("✅ Database migrated successfully!")
	}
}
