package productservice

import (
	"strings"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

// new Product service construtor
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

// generateSlug - Simple slug generation
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	replacer := strings.NewReplacer(
		".", "", ",", "", "!", "", "?", "", "'", "", "\"", "",
		"(", "", ")", "", "[", "", "]", "", "{", "", "}", "",
		"&", "and", "@", "at",
	)
	return replacer.Replace(slug)
}
