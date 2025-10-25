package cloudinarydto

// UploadImageOptions defines options for image upload
type UploadImageOptions struct {
	Folder           string   // Cloudinary folder path
	PublicID         string   // Custom public ID (optional)
	Transformation   string   // Image transformation (optional)
	AllowedFormats   []string // Allowed image formats
	MaxFileSizeBytes int64    // Maximum file size in bytes
	Tags             []string // Tags for organization
}
