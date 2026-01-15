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
	// Load .env file (optional for local development)
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using environment variables")
		}
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s client_encoding=UTF8",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	var err error
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
		&models.Admin{},
		&models.AdminInvitation{},
		&models.RefreshToken{},
		&models.Buyer{},
	)
	if err != nil {
		fmt.Println("Base tables migration failed:", err)
		return
	}

	// Step 2: Create dependent tables
	err = DB.AutoMigrate(
		&models.Product{},
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
		&models.SellerLoginActivity{},
		&models.SellerNotificationPreference{},
		&models.SellerPolicy{},
		&models.SellerVerificationDocument{},
		&models.StockHistory{},
		&models.InventoryThreshold{},
		&models.CommissionSetting{},
		&models.SellerEarningsSummary{},
		&models.SellerPayout{},
		&models.ReviewReport{},
		&models.OrderCommission{},
		&models.PlatformRevenueSummary{},
		&models.Message{},
		&models.Conversation{},
		&models.MessageAttachment{},
	)
	if err != nil {
		fmt.Println(" Migration failed:", err)
	} else {
		fmt.Println(" Database migrated successfully!")
	}

}
