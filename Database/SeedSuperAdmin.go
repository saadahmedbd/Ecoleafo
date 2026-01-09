package database

import (
	"log"
	"os"

	models "github.com/saadahmedbd/Treestore/Models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedSuperAdminFromEnv(db *gorm.DB) error {
	// Check if environment variables are set
	username := os.Getenv("SUPER_ADMIN_USERNAME")
	email := os.Getenv("SUPER_ADMIN_EMAIL")
	password := os.Getenv("SUPER_ADMIN_PASSWORD")
	fullName := os.Getenv("SUPER_ADMIN_FULLNAME")

	if username == "" || email == "" || password == "" {
		log.Println("Super admin environment variables not set, skipping...")
		return nil

	}
	// Check if RegUser with this email already exists
	var existingUser models.RegUser
	if err := db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		// User exists, check if admin profile exists
		var existingAdmin models.Admin
		if err := db.Where("user_id = ?", existingUser.ID).First(&existingAdmin).Error; err == nil {
			log.Println("Super admin already exists")
			return nil
		}
		// User exists but no admin profile, create admin profile
		if fullName == "" {
			fullName = "Super Administrator"
		}
		admin := &models.Admin{
			UserID:            existingUser.ID,
			Role:              "super_admin",
			FullName:          fullName,
			Department:        "IT",
			CanManageUsers:    true,
			CanManageProducts: true,
			CanManageOrders:   true,
			CanViewReports:    true,
			CanManageSettings: true,
			IsActive:          true,
		}
		if err := db.Create(admin).Error; err != nil {
			return err
		}
		log.Println("✅ Super admin profile created for existing user!")
		return nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Create user
	user := &models.RegUser{
		Email:    email,
		Password: string(hashedPassword),
		Role:     "admin",
		IsActive: true,
	}

	if err := db.Create(user).Error; err != nil {
		return err
	}

	if fullName == "" {
		fullName = "Super Administrator"
	}

	// Create admin profile
	admin := &models.Admin{
		UserID:            user.ID,
		Role:              "super_admin",
		FullName:          fullName,
		Department:        "IT",
		CanManageUsers:    true,
		CanManageProducts: true,
		CanManageOrders:   true,
		CanViewReports:    true,
		CanManageSettings: true,
		IsActive:          true,
	}

	if err := db.Create(admin).Error; err != nil {
		return err
	}
	log.Println("✅ Super admin created from environment variables!")
	return nil

}
