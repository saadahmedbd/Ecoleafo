package cloudniaryservice

import (
	"context"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	cloudinarydto "github.com/saadahmedbd/Treestore/Rest/DTO/CloudinaryDTO"
)

// CloudinaryConfig holds Cloudinary configuration
type CloudinaryConfig struct {
	CloudName string
	APIKey    string
	APISecret string
}

type CloudniaryService struct {
	cld *cloudinary.Cloudinary
}

// NewCloudinaryService creates a new Cloudinary service instance
func NewCloudinaryService(config CloudinaryConfig) *CloudniaryService {
	cld, err := cloudinary.NewFromParams(config.CloudName, config.APIKey, config.APISecret)
	if err != nil {
		return nil
	}

	return &CloudniaryService{cld: cld}
}

// UploadImage uploads an image to Cloudinary with validation
func (s *CloudniaryService) UploadImage(ctx context.Context, file io.Reader, filename string, opts cloudinarydto.UploadImageOptions) (*cloudinarydto.UploadResult, error) {
	// Validate file extension
	ext := strings.ToLower(filepath.Ext(filename))
	if !s.isValidImageFormat(ext, opts.AllowedFormats) {
		return nil, fmt.Errorf("invalid file format: %s", ext)
	}

	// Set default folder if not provided
	folder := opts.Folder
	if folder == "" {
		folder = "uploads"
	}

	// Prepare upload parameters
	uploadParams := uploader.UploadParams{
		Folder:       folder,
		ResourceType: "image",
		// Invalidate:   false, // Set to false by default to reduce costs
	}

	// Set public ID if provided
	if opts.PublicID != "" {
		uploadParams.PublicID = opts.PublicID
	}

	// Add tags if provided
	if len(opts.Tags) > 0 {
		uploadParams.Tags = opts.Tags
	}

	// Upload to Cloudinary
	result, err := s.cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return nil, fmt.Errorf("cloudinary upload failed: %w", err)
	}
	log.Printf("✅ Image uploaded successfully!\nPublicID: %s\nURL: %s\nBytes: %d\n",
		result.PublicID, result.SecureURL, result.Bytes)

	return &cloudinarydto.UploadResult{
		PublicID:  result.PublicID,
		URL:       result.SecureURL, // Prefer SecureURL for client responses
		SecureURL: result.SecureURL,
		Format:    result.Format,
		Width:     result.Width,
		Height:    result.Height,
		Bytes:     result.Bytes,
	}, nil
}

// DeleteImage deletes an image from Cloudinary
func (s *CloudniaryService) DeleteImage(ctx context.Context, publicID string) error {
	_, err := s.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",
		// Invalidate:   true,
	})
	if err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}
	return nil
}

// isValidImageFormat checks if the file format is allowed
func (s *CloudniaryService) isValidImageFormat(ext string, allowedFormats []string) bool {
	if len(allowedFormats) == 0 {
		// Default allowed formats
		allowedFormats = []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp"}
	}

	for _, format := range allowedFormats {
		if ext == format {
			return true
		}
	}
	return false
}
