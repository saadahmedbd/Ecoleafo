package adminmangement

import (
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{db: db}
}

// GetAll retrieves all sellers with pagination
func (r *SellerRepository) GetAll(page, limit int, status string) ([]models.User, int64, error) {
	var sellers []models.User
	var total int64

	query := r.db.Model(&models.User{}).Preload("RegUser")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&sellers).Error

	return sellers, total, err
}

// GetByID retrieves a seller by ID
func (r *SellerRepository) GetByID(id uint) (*models.User, error) {
	var seller models.User
	err := r.db.Preload("RegUser").Preload("ApprovedByAdmin").First(&seller, id).Error
	return &seller, err
}

// GetPendingApprovals retrieves sellers pending approval
func (r *SellerRepository) GetPendingApprovals(page, limit int) ([]models.User, int64, error) {
	var sellers []models.User
	var total int64

	query := r.db.Model(&models.User{}).
		Where("approval_status = ?", "pending").
		Preload("RegUser")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at ASC").Find(&sellers).Error

	return sellers, total, err
}

// Update updates a seller
func (r *SellerRepository) Update(seller *models.User) error {
	return r.db.Save(seller).Error
}

// UpdateStatus updates seller status
func (r *SellerRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("status", status).Error
}

// ApproveReject approves or rejects a seller
func (r *SellerRepository) ApproveReject(sellerID uint, status, reason string, adminUserID uint) error {
	var admin models.Admin

	// Find admin by user_id (from JWT)
	if err := r.db.Where("user_id = ?", adminUserID).First(&admin).Error; err != nil {
		return fmt.Errorf("admin not found for user_id %d: %w", adminUserID, err)
	}

	updates := map[string]interface{}{
		"approval_status": status,
		"status":          status,
		"approved_by":     admin.ID, // use actual admin.id
	}

	if reason != "" && status == "rejected" {
		updates["rejection_reason"] = reason
		updates["rejected_at"] = gorm.Expr("NOW()") // set rejection time
	}

	if status == "approved" {
		updates["approved_at"] = gorm.Expr("NOW()") // set approval time
	}

	result := r.db.Model(&models.User{}).Where("id = ?", sellerID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("seller with id %d not found", sellerID)
	}

	return nil
}

func (r *SellerRepository) ApproveSeller(sellerID uint, adminUserID uint) error {
	var admin models.Admin

	//  Fix: find admin by user_id (from JWT), not by admin.id
	if err := r.db.Where("user_id = ?", adminUserID).First(&admin).Error; err != nil {
		return fmt.Errorf("admin not found for user_id %d: %w", adminUserID, err)
	}

	// Update seller approval fields
	result := r.db.Model(&models.User{}).
		Where("id = ?", sellerID).
		Updates(map[string]interface{}{
			"approval_status": "approved",
			"approved_at":     time.Now(),
			"approved_by":     admin.ID, // store the actual admin.id here
			"status":          "approved",
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("seller with id %d not found", sellerID)
	}

	return nil
}

// Search searches sellers by business name or email
func (r *SellerRepository) Search(query string, page, limit int) ([]*models.User, int64, error) {
	var sellers []*models.User
	var total int64

	searchQuery := "%" + query + "%"

	// Count query
	countQuery := r.db.Model(&models.User{}).
		Joins("LEFT JOIN reg_users ON reg_users.id = users.user_id").
		Where("users.store_name ILIKE ? OR reg_users.email ILIKE ?", searchQuery, searchQuery)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	// Main query
	err := r.db.Model(&models.User{}).
		Joins("LEFT JOIN reg_users ON reg_users.id = users.user_id").
		Where("users.store_name ILIKE ? OR reg_users.email ILIKE ?", searchQuery, searchQuery).
		Preload("RegUser").
		Offset(offset).
		Limit(limit).
		Order("users.created_at DESC").
		Find(&sellers).Error

	if err != nil {
		return nil, 0, err
	}

	return sellers, total, nil
}

// GetStats retrieves seller statistics
func (r *SellerRepository) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var total, pending, approved, rejected, suspended int64
	r.db.Model(&models.User{}).Count(&total)
	r.db.Model(&models.User{}).Where("approval_status = ?", "pending").Count(&pending)
	r.db.Model(&models.User{}).Where("status = ?", "approved").Count(&approved)
	r.db.Model(&models.User{}).Where("status = ?", "rejected").Count(&rejected)
	r.db.Model(&models.User{}).Where("status = ?", "suspended").Count(&suspended)

	stats["total"] = total
	stats["pending"] = pending
	stats["approved"] = approved
	stats["rejected"] = rejected
	stats["suspended"] = suspended

	return stats, nil
}

// GetTopSellers retrieves top sellers by sales
func (r *SellerRepository) GetTopSellers(limit int) ([]models.User, error) {
	var sellers []models.User
	err := r.db.Preload("RegUser").
		Where("status = ?", "approved").
		Order("total_sales DESC").
		Limit(limit).
		Find(&sellers).Error
	return sellers, err
}
