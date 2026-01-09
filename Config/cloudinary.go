package Config

// ==========================================
// CLOUDINARY CONFIGURATION
// ==========================================
// Location: backend/config/cloudinary.go
// Purpose: Initialize Cloudinary client with credentials
// ==========================================

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

var CloudinaryClient *cloudinary.Cloudinary

// InitCloudinary initializes Cloudinary configuration
func InitCloudinary() error {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Get Cloudinary credentials from environment
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	// Validate credentials
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return fmt.Errorf("cloudinary credentials not set in environment variables")
	}

	// Create Cloudinary instance
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return fmt.Errorf("failed to initialize cloudinary: %w", err)
	}

	CloudinaryClient = cld
	log.Println("✓ Cloudinary initialized successfully")
	// fmt.Println("cloudinary config:", Config.CloudinaryClient.Config.Cloud)

	return nil
}

// GetCloudinaryClient returns the Cloudinary client instance
func GetCloudinaryClient() *cloudinary.Cloudinary {
	if CloudinaryClient == nil {
		log.Fatal("Cloudinary client not initialized. Call InitCloudinary() first")
	}
	return CloudinaryClient
}
