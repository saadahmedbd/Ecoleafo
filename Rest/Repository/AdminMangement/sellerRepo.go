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
	return r.db.Transaction(func(tx *gorm.DB) error {
		var admin models.Admin
		if err := tx.Where("user_id = ?", adminUserID).First(&admin).Error; err != nil {
			return fmt.Errorf("admin not found for user_id %d: %w", adminUserID, err)
		}

		var seller models.User
		if err := tx.First(&seller, sellerID).Error; err != nil {
			return fmt.Errorf("seller with id %d not found: %w", sellerID, err)
		}

		now := time.Now()

		//  Ensure basic business info is filled
		updateData := map[string]interface{}{
			"approval_status":     "approved",
			"approved_at":         now,
			"approved_by":         admin.ID,
			"status":              "approved",
			"is_approved":         true,
			"is_verified":         true,
			"is_active":           true,
			"is_profile_complete": true,
			"has_business_info":   true,
			"has_address":         true,
			"has_payment_method":  true,
			"can_add_products":    true,
			"missing_fields":      nil,
			"next_step":           "approved",
		}

		// 🧠 Auto-fill missing business info if needed
		if seller.BusinessEmail == "" || seller.BusinessEmail == "N/A" {
			updateData["business_email"] = fmt.Sprintf("%s@auto-verified.com", seller.StoreSlug)
		}
		if seller.Address == "" || seller.Address == "N/A" {
			updateData["address"] = "Auto Verified Address"
			updateData["city"] = "Dhaka"
			updateData["state"] = "Dhaka"
			updateData["postal_code"] = "1200"
		}
		if seller.Phone == "" {
			updateData["phone"] = "01700000000"
		}

		// ✅ Update seller record
		if err := tx.Model(&models.User{}).Where("id = ?", sellerID).Updates(updateData).Error; err != nil {
			return fmt.Errorf("failed to update seller: %w", err)
		}

		// ✅ Update linked RegUser as verified
		if err := tx.Model(&models.RegUser{}).Where("id = ?", seller.UserId).Updates(map[string]interface{}{
			"is_verified": true,
			"is_active":   true,
		}).Error; err != nil {
			return fmt.Errorf("failed to update RegUser: %w", err)
		}

		//  Optionally ensure at least one payment method exists
		var paymentCount int64
		tx.Model(&models.SellerPaymentMethod{}).Where("seller_id = ?", sellerID).Count(&paymentCount)
		if paymentCount == 0 {
			defaultMethod := models.SellerPaymentMethod{
				SellerID: seller.ID,
				// Method:    "Bank Transfer",
				// AccountNo: "000000000",
				IsActive: true,
			}
			if err := tx.Create(&defaultMethod).Error; err != nil {
				return fmt.Errorf("failed to add default payment method: %w", err)
			}
		}

		fmt.Printf(" Seller (UserID=%d, RegUserID=%d) approved and profile completed by AdminID=%d\n",
			sellerID, seller.UserId, adminUserID)

		return nil
	})
}
func (r *SellerRepository) UpdateSellerFields(id uint, fields map[string]interface{}) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(fields).Error
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
