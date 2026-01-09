package adminmangement

import (
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAll retrieves all products with pagination
func (r *ProductRepository) GetAll(page, limit int, status string) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.db.Model(&models.Product{}).
		Preload("Seller").
		Preload("Seller.RegUser").
		Preload("Category").
		Preload("Images")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&products).Error

	return products, total, err
}

// GetByID retrieves a product by ID
func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Seller").
		Preload("Seller.RegUser").
		Preload("Category").
		Preload("Images").
		Preload("ApprovedByAdmin").
		First(&product, id).Error
	return &product, err
}

// GetBySeller retrieves products by seller ID
func (r *ProductRepository) GetBySeller(sellerID uint, page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.db.Model(&models.Product{}).
		Where("seller_id = ?", sellerID).
		Preload("Category").
		Preload("Images")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&products).Error

	return products, total, err
}

// GetPendingApprovals retrieves products pending approval
func (r *ProductRepository) GetPendingApprovals(page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.db.Model(&models.Product{}).
		Where("approval_status = ?", "pending").
		Preload("Seller").
		Preload("Seller.RegUser").
		Preload("Category").
		Preload("Images")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at ASC").Find(&products).Error

	return products, total, err
}

// Create creates a new product
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// Update updates a product
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// ApproveReject approves or rejects a product
func (r *ProductRepository) ApproveProduct(productID uint, adminUserID uint) error {
	var admin models.Admin

	//  Fix: find admin by user_id (from JWT), not by admin.id
	if err := r.db.Where("user_id = ?", adminUserID).First(&admin).Error; err != nil {
		return fmt.Errorf("admin not found for user_id %d: %w", adminUserID, err)
	}

	// Update seller approval fields
	result := r.db.Model(&models.Product{}).
		Where("id = ?", productID).
		Updates(map[string]interface{}{
			"approval_status": "approved",
			"approved_at":     time.Now(),
			"approved_by":     admin.ID, // store the actual admin.id here

		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product with id %d not found", productID)
	}

	return nil
}

func (r *ProductRepository) ApproveReject(productID uint, status, reason string, adminUserID uint) error {
	var admin models.Admin

	if err := r.db.Where("user_id = ?", adminUserID).First(&admin).Error; err != nil {
		return fmt.Errorf("admin not found for user_id %d: %w", adminUserID, err)
	}

	updates := map[string]interface{}{
		"approval_status": status,
		"approved_by":     admin.ID,
	}

	if status == "rejected" {
		updates["approved_at"] = nil
		updates["rejection_reason"] = reason
	} else if status == "approved" {
		now := time.Now()
		updates["approved_at"] = &now
		updates["rejection_reason"] = ""
	}

	result := r.db.Model(&models.Product{}).Where("id = ?", productID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("product with id %d not found", productID)
	}

	return nil
}

// UpdateStock updates product stock quantity
func (r *ProductRepository) UpdateStock(id uint, quantity int) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).Update("stock_quantity", quantity).Error
}

// Search searches products by name or SKU
func (r *ProductRepository) Search(query string, page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	searchQuery := "%" + query + "%"
	dbQuery := r.db.Model(&models.Product{}).
		Where("name ILIKE ? OR sku ILIKE ?", searchQuery, searchQuery).
		Preload("Seller").
		Preload("Seller.RegUser").
		Preload("Category").
		Preload("Images")

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := dbQuery.Offset(offset).Limit(limit).Order("created_at DESC").Find(&products).Error

	return products, total, err
}

// Delete soft deletes a product
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

// GetStats retrieves product statistics
func (r *ProductRepository) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var total, pending, approved, rejected, lowStock int64
	r.db.Model(&models.Product{}).Count(&total)
	r.db.Model(&models.Product{}).Where("approval_status = ?", "pending").Count(&pending)
	r.db.Model(&models.Product{}).Where("approval_status = ?", "approved").Count(&approved)
	r.db.Model(&models.Product{}).Where("approval_status = ?", "rejected").Count(&rejected)
	r.db.Model(&models.Product{}).Where("stock_quantity < min_stock_level").Count(&lowStock)

	stats["total"] = total
	stats["pending"] = pending
	stats["approved"] = approved
	stats["rejected"] = rejected
	stats["low_stock"] = lowStock

	return stats, nil
}

// GetTopProducts retrieves top products by sale count
func (r *ProductRepository) GetTopProducts(limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("approval_status = ?", "approved").
		Preload("Seller").
		Preload("Seller.RegUser").
		Preload("Category").
		Preload("Images").
		Order("sale_count DESC").
		Limit(limit).
		Find(&products).Error
	return products, err
}
