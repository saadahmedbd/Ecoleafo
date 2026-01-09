package repository

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"

	"gorm.io/gorm"
)

type SellerRepository interface {
	GetSellerByUserID(userID uint) (*models.User, error)
	GetSellerByID(sellerID uint) (*models.User, error)
	UpdateSellerProfile(seller *models.User) error
	UpdateRegUser(regUser *models.RegUser) error //for changing password
	GetSellerStats(sellerID uint) (*sellerprofile.SellerStateResponse, error)
	GetSellerWithRelations(sellerID uint) (*models.User, error)
}

type sellerRepository struct {
	db *gorm.DB
}

func NewSellerRepostory(db *gorm.DB) SellerRepository {
	return &sellerRepository{
		db: db,
	}
}
func (r *sellerRepository) GetSellerByUserID(userID uint) (*models.User, error) {
	var seller models.User
	err := r.db.Where("user_id = ?", userID).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller not found")
		}
		return &seller, fmt.Errorf("failed to get seller: %v", err)
	}
	return &seller, nil
}

func (r *sellerRepository) GetSellerByID(sellerID uint) (*models.User, error) {
	var seller models.User
	err := r.db.Where("id = ?", sellerID).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller not found")
		}
		return nil, fmt.Errorf("failed to fetch seller: %v", err)
	}
	return &seller, nil
}

func (r *sellerRepository) GetSellerWithRelations(sellerID uint) (*models.User, error) {
	var seller models.User
	err := r.db.
		Preload("RegUser").
		Preload("Role").
		Preload("Products", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, seller_id, name, price, quantity, is_active, is_approved")
		}).
		Where("id = ?", sellerID).
		First(&seller).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("seller not found")
		}
		return nil, fmt.Errorf("failed to fetch seller: %v", err)
	}
	return &seller, nil
}
func (r *sellerRepository) UpdateSellerProfile(seller *models.User) error {
	return r.db.Save(seller).Error
}

func (r *sellerRepository) UpdateRegUser(regUser *models.RegUser) error {
	return r.db.Save(regUser).Error
}

func (r *sellerRepository) GetSellerStats(sellerID uint) (*sellerprofile.SellerStateResponse, error) {
	stats := &sellerprofile.SellerStateResponse{}

	// Get basic seller info
	var seller models.User
	if err := r.db.Where("id = ?", sellerID).First(&seller).Error; err != nil {
		return nil, err
	}

	stats.TotalSales = seller.TotalSales
	stats.TotalEarnings = seller.TotalEarnings
	stats.TotalOrders = seller.TotalOrders
	stats.AverageRating = seller.AverageRating

	// Count products by status - using int64 variables for GORM Count
	var totalProducts, activeProducts, pendingProducts, reviewCount int64

	r.db.Model(&models.Product{}).
		Where("seller_id = ?", sellerID).
		Count(&totalProducts)
	stats.TotalProducts = int(totalProducts)

	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND is_active = ?", sellerID, true).
		Count(&activeProducts)
	stats.ActiveProducts = int(activeProducts)

	r.db.Model(&models.Product{}).
		Where("seller_id = ? AND is_approved = ?", sellerID, false).
		Count(&pendingProducts)
	stats.PendingProducts = int(pendingProducts)

	// Count reviews
	r.db.Model(&models.Review{}).
		Joins("JOIN products ON reviews.product_id = products.id").
		Where("products.seller_id = ?", sellerID).
		Count(&reviewCount)
	stats.ReviewCount = int(reviewCount)

	return stats, nil
}
