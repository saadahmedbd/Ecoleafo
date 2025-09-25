package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	database "github.com/saadahmedbd/Treestore/Database"
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase connects to PostgreSQL using GORM
func Connect() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connectd to postgres database successfully")

	// Clean up reviews table to fix foreign key constraint
	DB.Exec("DELETE FROM reviews WHERE order_id = 0 OR order_id NOT IN (SELECT id FROM orders)")

	// Run Auto Migration
	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderItem{},
		&models.Buyer{},
		&models.CartItem{},
		&models.Review{},
		&models.Role{},
		&models.Wishlist{},
		&models.Address{},
		&models.AuditLog{},
		&models.ReviewImage{},
		&models.ProductAttribute{},
		&models.ProductImage{},
		&models.SellerPaymentMethod{},
		&models.Setting{},
		&models.SellerCategory{},
		&models.OrderHistory{},
		&models.Notification{},
	)
	if err != nil {
		fmt.Println(" Migration failed:", err)
	} else {
		fmt.Println(" Database migrated successfully!")
	}

	// Migrate RegUser separately
	if err := database.MigrateRegUser(DB); err != nil {
		fmt.Println(" RegUser migration failed:", err)
	}

	// Seed data

	// if err := database.SeedData(DB); err != nil {
	// 	log.Fatal(" Seeding failed:", err)
	// }

}
