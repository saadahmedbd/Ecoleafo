package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

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

	fmt.Println("Connected to postgres database successfully")

	// Run Auto Migration in proper order (parent tables first)
	// Step 1: Create base tables without foreign key constraints
	err = DB.AutoMigrate(
		&models.Role{},
		&models.RegUser{},
		&models.User{},
		&models.Category{},
	)
	if err != nil {
		fmt.Println("Base tables migration failed:", err)
		return
	}

	// Step 2: Create dependent tables
	err = DB.AutoMigrate(
		&models.Product{},
		&models.Buyer{},
		&models.Order{},
		&models.OrderItem{},
		&models.CartItem{},
		&models.GuestCartItem{},
		&models.Review{},
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

}
