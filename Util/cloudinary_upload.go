package util

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/saadahmedbd/Treestore/Config"
)

// FileUploadConfig holds upload configuration
type FileUploadConfig struct {
	Folder         string   // Cloudinary folder (e.g., "seller-photos")
	AllowedFormats []string // Allowed file formats (e.g., ["jpg", "png"])
	MaxSizeBytes   int64    // Maximum file size in bytes
	Width          int      // Image width (0 for original)
	Height         int      // Image height (0 for original)
	Quality        string   // Image quality: "auto", "best", or 1-100
	Format         string   // Output format: "jpg", "png", "webp", etc.
	Transformation string   // Custom transformation string
}

// DefaultImageConfig returns default image upload configuration
func DefaultImageConfig(folder string) FileUploadConfig {
	return FileUploadConfig{
		Folder:         folder,
		AllowedFormats: []string{".jpg", ".jpeg", ".png", ".gif", ".webp"},
		MaxSizeBytes:   5 * 1024 * 1024, // 5MB
		Quality:        "auto",
		Format:         "auto",
	}
}

// UploadImageToCloudinary uploads an image to Cloudinary
func UploadImageToCloudinary(file multipart.File, header *multipart.FileHeader, cfg FileUploadConfig) (string, error) {
	// Validate file size
	if header.Size > cfg.MaxSizeBytes {
		return "", fmt.Errorf("file size exceeds maximum allowed size of %d bytes", cfg.MaxSizeBytes)
	}

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !isAllowedFormat(ext, cfg.AllowedFormats) {
		return "", fmt.Errorf("file format %s not allowed. Allowed formats: %v", ext, cfg.AllowedFormats)
	}

	// Get Cloudinary client
	cld := Config.GetCloudinaryClient()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Generate unique public ID
	publicID := generatePublicID(header.Filename)

	// Prepare upload parameters
	uploadParams := uploader.UploadParams{
		PublicID:       publicID,
		Folder:         cfg.Folder,
		ResourceType:   "image",
		Overwrite:      boolPtr(false),
		UniqueFilename: boolPtr(true),
		UseFilename:    boolPtr(true),
	}

	// Add transformation parameters if specified
	if cfg.Width > 0 || cfg.Height > 0 || cfg.Quality != "" || cfg.Format != "" {
		transformation := buildTransformation(cfg)
		uploadParams.Transformation = transformation
	}

	// Upload file to Cloudinary
	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		fmt.Println("cloudinary upload error:", err)
		return "", fmt.Errorf("failed to upload file to cloudinary: %w", err)
	}
	fmt.Printf("✅ Cloudinary upload result: %+v\n", result)

	if result.SecureURL == "" {
		return "", fmt.Errorf("cloudinary returned empty secure URL (check API credentials or upload params)")
	}
	// Return the secure URL
	return result.SecureURL, nil
}

// UploadImageToCloudinaryWithPublicID uploads an image to Cloudinary and returns URL and public ID
func UploadImageToCloudinaryWithPublicID(file multipart.File, header *multipart.FileHeader, cfg FileUploadConfig) (string, string, error) {
	// Validate file size
	if header.Size > cfg.MaxSizeBytes {
		return "", "", fmt.Errorf("file size exceeds maximum allowed size of %d bytes", cfg.MaxSizeBytes)
	}

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !isAllowedFormat(ext, cfg.AllowedFormats) {
		return "", "", fmt.Errorf("file format %s not allowed. Allowed formats: %v", ext, cfg.AllowedFormats)
	}

	// Get Cloudinary client
	cld := Config.GetCloudinaryClient()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Generate unique public ID
	publicID := generatePublicID(header.Filename)

	// Prepare upload parameters
	uploadParams := uploader.UploadParams{
		PublicID:       publicID,
		Folder:         cfg.Folder,
		ResourceType:   "image",
		Overwrite:      boolPtr(false),
		UniqueFilename: boolPtr(true),
		UseFilename:    boolPtr(true),
	}

	// Add transformation parameters if specified
	if cfg.Width > 0 || cfg.Height > 0 || cfg.Quality != "" || cfg.Format != "" {
		transformation := buildTransformation(cfg)
		uploadParams.Transformation = transformation
	}

	// Upload file to Cloudinary
	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		fmt.Println("cloudinary upload error:", err)
		return "", "", fmt.Errorf("failed to upload file to cloudinary: %w", err)
	}
	fmt.Printf("✅ Cloudinary upload result: %+v\n", result)

	if result.SecureURL == "" {
		return "", "", fmt.Errorf("cloudinary returned empty secure URL (check API credentials or upload params)")
	}

	// Construct full public ID including folder
	fullPublicID := cfg.Folder + "/" + publicID

	// Return the secure URL and public ID
	return result.SecureURL, fullPublicID, nil
}

// UploadProfilePhoto uploads a profile photo with optimizations
func UploadProfilePhoto(file multipart.File, header *multipart.FileHeader) (string, error) {
	cfg := FileUploadConfig{
		Folder:         "seller-profiles",
		AllowedFormats: []string{".jpg", ".jpeg", ".png", ".webp"},
		MaxSizeBytes:   5 * 1024 * 1024, // 2MB
		Width:          400,             // Resize to 400x400
		Height:         400,
		Quality:        "auto:good",
		Format:         "jpg",
	}
	return UploadImageToCloudinary(file, header, cfg)
}

// UploadBuyerProfilePhoto uploads a buyer profile photo with optimizations
func UploadBuyerProfilePhoto(file multipart.File, header *multipart.FileHeader) (string, string, error) {
	cfg := FileUploadConfig{
		Folder:         "buyer-profiles",
		AllowedFormats: []string{".jpg", ".jpeg", ".png", ".webp"},
		MaxSizeBytes:   5 * 1024 * 1024, // 5MB
		Height:         400,
		Quality:        "auto:good",
		Format:         "jpg",
	}
	return UploadImageToCloudinaryWithPublicID(file, header, cfg)
}

// UploadStoreLogo uploads a store logo
func UploadStoreLogo(file multipart.File, header *multipart.FileHeader) (string, error) {
	cfg := FileUploadConfig{
		Folder:         "store-logos",
		AllowedFormats: []string{".jpg", ".jpeg", ".png"},
		MaxSizeBytes:   2 * 1024 * 1024, // 2MB
		Width:          300,
		Height:         300,
		Quality:        "auto:best",
		Format:         "png", // PNG for logos with transparency
	}
	return UploadImageToCloudinary(file, header, cfg)
}

// UploadStoreBanner uploads a store banner
func UploadStoreBanner(file multipart.File, header *multipart.FileHeader) (string, error) {
	cfg := FileUploadConfig{
		Folder:         "store-banners",
		AllowedFormats: []string{".jpg", ".jpeg", ".png"},
		MaxSizeBytes:   5 * 1024 * 1024, // 5MB
		Width:          1200,
		Height:         300,
		Quality:        "auto:good",
		Format:         "jpg",
	}
	return UploadImageToCloudinary(file, header, cfg)
}

// UploadVerificationDocument uploads verification documents
func UploadVerificationDocument(file multipart.File, header *multipart.FileHeader) (string, error) {
	cfg := FileUploadConfig{
		Folder:         "verification-documents",
		AllowedFormats: []string{".jpg", ".jpeg", ".png", ".pdf"},
		MaxSizeBytes:   5 * 1024 * 1024, // 5MB
		Quality:        "auto",
	}
	return UploadImageToCloudinary(file, header, cfg)
}

// UploadProductImage uploads a product image
func UploadProductImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	cfg := FileUploadConfig{
		Folder:         "products",
		AllowedFormats: []string{".jpg", ".jpeg", ".png", ".webp"},
		MaxSizeBytes:   5 * 1024 * 1024, // 5MB
		Width:          800,
		Height:         800,
		Quality:        "auto:good",
		Format:         "webp", // WebP for better compression
	}
	return UploadImageToCloudinary(file, header, cfg)
}

// DeleteImageFromCloudinary deletes an image from Cloudinary
func DeleteImageFromCloudinary(imageURL string) error {
	// Extract public ID from Cloudinary URL
	publicID := extractPublicIDFromURL(imageURL)
	if publicID == "" {
		return errors.New("invalid cloudinary URL")
	}

	// Get Cloudinary client
	cld := Config.GetCloudinaryClient()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete the image
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",
	})

	if err != nil {
		return fmt.Errorf("failed to delete image from cloudinary: %w", err)
	}

	return nil
}

// ==========================================
// HELPER FUNCTIONS
// ==========================================

// generatePublicID generates a unique public ID for the uploaded file
func generatePublicID(filename string) string {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	// Remove special characters
	name = strings.ReplaceAll(name, " ", "_")

	// Add timestamp for uniqueness
	timestamp := time.Now().Unix()

	return fmt.Sprintf("%s_%d", name, timestamp)
}

// buildTransformation builds Cloudinary transformation string
func buildTransformation(cfg FileUploadConfig) string {
	var transformations []string

	if cfg.Width > 0 && cfg.Height > 0 {
		transformations = append(transformations, fmt.Sprintf("c_fill,w_%d,h_%d", cfg.Width, cfg.Height))
	} else if cfg.Width > 0 {
		transformations = append(transformations, fmt.Sprintf("c_scale,w_%d", cfg.Width))
	} else if cfg.Height > 0 {
		transformations = append(transformations, fmt.Sprintf("c_scale,h_%d", cfg.Height))
	}

	if cfg.Quality != "" && cfg.Quality != "auto" {
		transformations = append(transformations, fmt.Sprintf("q_%s", cfg.Quality))
	}

	if cfg.Format != "" && cfg.Format != "auto" {
		transformations = append(transformations, fmt.Sprintf("f_%s", cfg.Format))
	}

	return strings.Join(transformations, ",")
}

// extractPublicIDFromURL extracts public ID from Cloudinary URL
func extractPublicIDFromURL(url string) string {
	// Example URL: https://res.cloudinary.com/demo/image/upload/v1234567890/folder/filename.jpg
	parts := strings.Split(url, "/upload/")
	if len(parts) != 2 {
		return ""
	}

	// Get the part after /upload/
	pathParts := strings.Split(parts[1], "/")
	if len(pathParts) < 2 {
		return ""
	}

	// Skip version (v1234567890) and get folder/filename
	publicIDParts := pathParts[1:]
	publicID := strings.Join(publicIDParts, "/")

	// Remove file extension
	lastDot := strings.LastIndex(publicID, ".")
	if lastDot > 0 {
		publicID = publicID[:lastDot]
	}

	return publicID
}

// isAllowedFormat checks if the file format is allowed
func isAllowedFormat(ext string, allowedFormats []string) bool {
	for _, format := range allowedFormats {
		if strings.EqualFold(ext, format) {
			return true
		}
	}
	return false
}

// boolPtr returns a pointer to a boolean value
func boolPtr(b bool) *bool {
	return &b
}

// ==========================================
// ADVANCED FEATURES
// ==========================================

// GetOptimizedImageURL returns an optimized version of the image URL
func GetOptimizedImageURL(originalURL string, width, height int, quality string) string {
	// Extract public ID
	publicID := extractPublicIDFromURL(originalURL)
	if publicID == "" {
		return originalURL
	}

	// Get cloud name from config
	cloudName := Config.CloudinaryClient.Config.Cloud.CloudName

	// Build optimized URL
	transformation := fmt.Sprintf("c_fill,w_%d,h_%d,q_%s", width, height, quality)
	optimizedURL := fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s/%s",
		cloudName, transformation, publicID)

	return optimizedURL
}

// GetThumbnailURL returns a thumbnail version of the image
func GetThumbnailURL(originalURL string) string {
	return GetOptimizedImageURL(originalURL, 150, 150, "auto")
}
