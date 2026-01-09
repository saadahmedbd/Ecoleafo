package productservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// CreateProductRequest - DTO for creating products (No seller_id needed - auto-detected)
type CreateProductRequest struct {
	Name            string                    `json:"name" validate:"required,min=1,max=255"`
	Description     string                    `json:"description"`
	SKU             string                    `json:"sku" validate:"required,min=1,max=100"`
	CategoryID      uint                      `json:"category_id" validate:"required"`
	Price           float64                   `json:"price" validate:"required,gt=0"`
	DiscountPrice   float64                   `json:"discount_price"`
	DiscountPercent float64                   `json:"discount_percent"`
	Height          string                    `json:"height"`
	Age             string                    `json:"age"`
	TreeType        string                    `json:"tree_type"`
	PotSize         string                    `json:"pot_size"`
	ScientificName  string                    `json:"scientific_name"`
	CommonNames     string                    `json:"common_names"`
	Quantity        int                       `json:"quantity"`
	MinQuantity     int                       `json:"min_quantity"`
	Weight          float64                   `json:"weight"`
	MetaTitle       string                    `json:"meta_title"`
	MetaDescription string                    `json:"meta_description"`
	Images          []ProductImageRequest     `json:"images"`     // Multiple images support
	Attributes      []ProductAttributeRequest `json:"attributes"` // Custom attributes
}

// ProductImageRequest - For image uploads
type ProductImageRequest struct {
	ImageURL  string `json:"image_url" validate:"required"`
	AltText   string `json:"alt_text"`
	IsPrimary bool   `json:"is_primary"`
	SortOrder int    `json:"sort_order"`
}

// ProductAttributeRequest - For custom attributes
type ProductAttributeRequest struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

func (s *ProductService) CreateProduct(req CreateProductRequest, userIdFromJWT uint) (*models.Product, error) {
	// start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// Generate slug from name
	slug := generateSlug(req.Name)
	// Check if slug already exists and make it unique
	var count int64
	s.db.Model(&models.Product{}).Where("slug LIKE ?", slug+"%").Count(&count)
	if count > 0 {
		slug = fmt.Sprintf("%s-%d", slug, count+1)
	}
	// Find seller by user_id (not id) since JWT contains user_id
	var seller models.User
	if err := s.db.Where("user_id = ?", userIdFromJWT).First(&seller).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller with user_id %d does not exist", userIdFromJWT)
		}
		return nil, fmt.Errorf("failed to verify seller: %v", err)
	}

	product := models.Product{
		SellerID:        seller.ID, // Automatically set from JWT token
		Name:            req.Name,
		Slug:            slug,
		Description:     req.Description,
		SKU:             req.SKU,
		CategoryID:      req.CategoryID,
		Price:           req.Price,
		DiscountPrice:   req.DiscountPrice,
		DiscountPercent: req.DiscountPercent,
		Height:          req.Height,
		Age:             req.Age,
		TreeType:        req.TreeType,
		PotSize:         req.PotSize,
		ScientificName:  req.ScientificName,
		CommonNames:     req.CommonNames,
		Quantity:        req.Quantity,
		MinQuantity:     req.MinQuantity,
		Weight:          req.Weight,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
	}
	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if len(req.Images) > 0 {
		for i, imgReq := range req.Images {
			image := models.ProductImage{
				ProductID: product.ID,
				ImageURL:  imgReq.ImageURL,
				AltText:   imgReq.AltText,
				IsPrimary: imgReq.IsPrimary,
				SortOrder: imgReq.SortOrder,
			}
			// If no sort order specified, use index
			if image.SortOrder == 0 {
				image.SortOrder = i + 1
			}
			if err := tx.Create(&image).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}
	// Add product attributes
	if len(req.Attributes) > 0 {
		for _, attrReq := range req.Attributes {
			attribute := models.ProductAttribute{
				ProductID: product.ID,
				Name:      attrReq.Name,
				Value:     attrReq.Value,
			}
			if err := tx.Create(&attribute).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}
	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	// Load the created product with relationships
	return s.GetProductWithRelations(product.ID)

}
