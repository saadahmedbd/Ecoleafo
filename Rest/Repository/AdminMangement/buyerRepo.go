package adminmangement

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type BuyerRepository struct {
	db *gorm.DB
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}

// GetAll retrieves all Buyers with pagination
func (r *BuyerRepository) GetAll(page, limit int, status string) ([]models.Buyer, int64, error) {
	var Buyers []models.Buyer
	var total int64

	query := r.db.Model(&models.Buyer{}).Preload("RegUser")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&Buyers).Error

	return Buyers, total, err
}

// GetByID retrieves a Buyer by ID
func (r *BuyerRepository) GetByID(id uint) (*models.Buyer, error) {
	var Buyer models.Buyer
	err := r.db.Preload("RegUser").Preload("Addresses").First(&Buyer, id).Error
	return &Buyer, err
}

// GetByregUserID retrieves a Buyer by regUser ID
func (r *BuyerRepository) GetByregUserID(regUserID uint) (*models.Buyer, error) {
	var Buyer models.Buyer
	err := r.db.Preload("RegUser").Where("reg_Buyer_id = ?", regUserID).First(&Buyer).Error
	return &Buyer, err
}

// Search searches Buyers by name or email
func (r *BuyerRepository) Search(query string, page, limit int) ([]models.Buyer, int64, error) {
	var buyers []models.Buyer
	var total int64

	searchQuery := "%" + query + "%"

	// Use actual table names in PostgreSQL
	dbQuery := r.db.Model(&models.Buyer{}).
		Joins("LEFT JOIN reg_users ON reg_users.id = buyers.user_id"). // correct table & column
		Where("reg_users.first_name ILIKE ? OR reg_users.last_name ILIKE ? OR reg_users.email ILIKE ?",
			searchQuery, searchQuery, searchQuery).
		Preload("RegUser")

	// Count total results
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Paginate
	offset := (page - 1) * limit
	err := dbQuery.Offset(offset).Limit(limit).Order("buyers.created_at DESC").Find(&buyers).Error

	return buyers, total, err
}

// Update updates a Buyer
func (r *BuyerRepository) Update(Buyer *models.Buyer) error {
	return r.db.Save(Buyer).Error
}

// UpdateStatus updates Buyer status
func (r *BuyerRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Buyer{}).Where("id = ?", id).Update("status", status).Error
}

// Delete soft deletes a Buyer
func (r *BuyerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Buyer{}, id).Error
}

// GetStats retrieves Buyer statistics
func (r *BuyerRepository) GetStats() (map[string]int64, error) {
	stats := make(map[string]int64)
	var total, active, inactive, suspended int64

	r.db.Model(&models.Buyer{}).Count(&total)
	r.db.Model(&models.Buyer{}).Where("status = ?", "active").Count(&active)
	r.db.Model(&models.Buyer{}).Where("status = ?", "inactive").Count(&inactive)
	r.db.Model(&models.Buyer{}).Where("status = ?", "suspended").Count(&suspended)

	stats["total"] = total
	stats["active"] = active
	stats["inactive"] = inactive
	stats["suspended"] = suspended

	return stats, nil
}
