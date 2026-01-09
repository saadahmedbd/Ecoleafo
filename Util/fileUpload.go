package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const (
	UploadDir   = "./uploads"
	MaxFileSize = 10 << 20 // 10MB
)

// UploadFile handles file upload and returns the file URL
func UploadFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	// Create upload directory if it doesn't exist
	uploadPath := filepath.Join(UploadDir, folder)
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), generateRandomString(10), ext)
	filePath := filepath.Join(uploadPath, filename)

	// Create the file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Return the file URL (adjust based on your server setup)
	fileURL := fmt.Sprintf("/uploads/%s/%s", folder, filename)
	return fileURL, nil
}

// generateRandomString generates a random string of given length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// DeleteFile deletes a file from the server
func DeleteFile(fileURL string) error {
	// Convert URL to file path
	filePath := filepath.Join(".", fileURL)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil // File doesn't exist, no error
	}

	// Delete the file
	return os.Remove(filePath)
}

// ValidateFileSize validates file size
func ValidateFileSize(size int64, maxSize int64) error {
	if size > maxSize {
		return fmt.Errorf("file size exceeds maximum allowed size of %d bytes", maxSize)
	}
	return nil
}

// ValidateFileType validates file extension
func ValidateFileType(filename string, allowedTypes []string) error {
	ext := filepath.Ext(filename)
	for _, allowed := range allowedTypes {
		if ext == allowed {
			return nil
		}
	}
	return fmt.Errorf("file type %s not allowed", ext)
}
