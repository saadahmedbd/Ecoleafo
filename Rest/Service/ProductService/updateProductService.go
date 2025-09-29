package productservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// UpdateProductRequest - DTO for updating products
type UpdateProductRequest struct {
	Name            *string                    `json:"name,omitempty"`
	Description     *string                    `json:"description,omitempty"`
	Price           *float64                   `json:"price,omitempty"`
	DiscountPrice   *float64                   `json:"discount_price,omitempty"`
	DiscountPercent *float64                   `json:"discount_percent,omitempty"`
	Height          *string                    `json:"height,omitempty"`
	Age             *string                    `json:"age,omitempty"`
	TreeType        *string                    `json:"tree_type,omitempty"`
	PotSize         *string                    `json:"pot_size,omitempty"`
	ScientificName  *string                    `json:"scientific_name,omitempty"`
	CommonNames     *string                    `json:"common_names,omitempty"`
	Quantity        *int                       `json:"quantity,omitempty"`
	MinQuantity     *int                       `json:"min_quantity,omitempty"`
	Weight          *float64                   `json:"weight,omitempty"`
	IsActive        *bool                      `json:"is_active,omitempty"`
	IsApproved      *bool                      `json:"is_approved,omitempty"`
	IsFeatured      *bool                      `json:"is_featured,omitempty"`
	MetaTitle       *string                    `json:"meta_title,omitempty"`
	MetaDescription *string                    `json:"meta_description,omitempty"`
	Images          *[]ProductImageRequest     `json:"images,omitempty"`
	Attributes      *[]ProductAttributeRequest `json:"attributes,omitempty"`
}

func (s *ProductService) UpdateProduct(ProductID uint, req UpdateProductRequest, userIdFromJWT uint) (*models.Product, error) {
	// Start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// fmt.Printf("DEBUG Service: productID=%d, userIdFromJWT=%d\n", ProductID, userIdFromJWT)
	// First, find the seller by user_id to get their actual ID
	var seller models.User
	if err := s.db.Where("user_id = ?", userIdFromJWT).First(&seller).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller with user_id %d does not exist", userIdFromJWT)
		}
		return nil, fmt.Errorf("failed to verify seller: %v", err)
	}

	// fmt.Printf("DEBUG Service: Found seller - ID=%d, user_id=%d, store=%s\n", seller.ID, seller.UserId, seller.StoreName)

	var product models.Product
	// Check if product belongs to seller
	if err := tx.Where("id = ? AND seller_id = ?", ProductID, seller.ID).First(&product).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product not found or you don't have permission to update it")
		}
		return nil, fmt.Errorf("product not found or access denied")
	}
	// fmt.Printf("DEBUG Service: Found product - ID=%d, seller_id=%d, name=%s\n", product.ID, product.SellerID, product.Name)

	// Update basic fields
	// Update product fields
	if req.Name != nil {
		product.Name = *req.Name
		// Regenerate slug if name changed
		slug := generateSlug(*req.Name)
		var count int64
		s.db.Model(&models.Product{}).Where("slug LIKE ? AND id != ?", slug+"%", ProductID).Count(&count)
		if count > 0 {
			slug = fmt.Sprintf("%s-%d", slug, count+1)
		}
		product.Slug = slug
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.DiscountPrice != nil {
		product.DiscountPrice = *req.DiscountPrice
	}
	if req.DiscountPercent != nil {
		product.DiscountPercent = *req.DiscountPercent
	}
	if req.Height != nil {
		product.Height = *req.Height
	}
	if req.Age != nil {
		product.Age = *req.Age
	}
	if req.TreeType != nil {
		product.TreeType = *req.TreeType
	}
	if req.PotSize != nil {
		product.PotSize = *req.PotSize
	}
	if req.ScientificName != nil {
		product.ScientificName = *req.ScientificName
	}
	if req.CommonNames != nil {
		product.CommonNames = *req.CommonNames
	}
	if req.Quantity != nil {
		product.Quantity = *req.Quantity
	}
	if req.MinQuantity != nil {
		product.MinQuantity = *req.MinQuantity
	}
	if req.Weight != nil {
		product.Weight = *req.Weight
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}
	if req.IsApproved != nil {
		product.IsApproved = *req.IsApproved
	}
	if req.IsFeatured != nil {
		product.IsFeatured = *req.IsFeatured
	}
	if req.MetaTitle != nil {
		product.MetaTitle = *req.MetaTitle
	}
	if req.MetaDescription != nil {
		product.MetaDescription = *req.MetaDescription
	}

	// Save basic product updates
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update images if provided
	if req.Images != nil {
		// Delete existing images
		if err := tx.Where("product_id = ?", product.ID).Delete(&models.ProductImage{}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Add new images
		for i, imgReq := range *req.Images {
			image := models.ProductImage{
				ProductID: product.ID,
				ImageURL:  imgReq.ImageURL,
				AltText:   imgReq.AltText,
				IsPrimary: imgReq.IsPrimary,
				SortOrder: imgReq.SortOrder,
			}
			if image.SortOrder == 0 {
				image.SortOrder = i + 1
			}
			if err := tx.Create(&image).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Update attributes if provided
	if req.Attributes != nil {
		// Delete existing attributes
		if err := tx.Where("product_id = ?", product.ID).Delete(&models.ProductAttribute{}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Add new attributes
		for _, attrReq := range *req.Attributes {
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

	// Return updated product with relationships
	return s.GetProductWithRelations(product.ID)
}
