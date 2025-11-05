package buyerProfilerepo

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"

	"gorm.io/gorm"
)

type BuyerRepository interface {
	// Define methods for buyer repository here
	GetBuyerByUserID(userID uint) (*models.Buyer, error)
	GetBuyerByID(buyerID uint) (*models.Buyer, error)
	UpdateBuyerProfile(buyer *models.Buyer) error
	GetBuyerWithRelations(buyerID uint) (*models.Buyer, error)
	GetBuyerStats(buyerID uint) (*buyerprofile.BuyerStateResponse, error)
	GetBuyerOrders(buyerID uint, page, limit int) ([]models.Order, int64, error)

	// Address management
	CreateAddress(address *models.Address) error
	UpdateAddress(address *models.Address) error
	DeleteAddress(addressID, buyerID uint) error
	GetAddressByID(addressID, buyerID uint) (*models.Address, error)
	GetBuyerAddresses(buyerID uint) ([]models.Address, error)
	SetDefaultAddress(buyerID, addressID uint) error
	UpdateRegUser(regUser *models.RegUser) error
	/// create udate reguser for ulpad profile photo\
	UpdateRegUserForImage(userID uint, updates map[string]interface{}) error
	ClearProfilePhoto(userID uint) error
}
type buyerRepository struct {
	db *gorm.DB
}

func NewBuyerRepository(db *gorm.DB) BuyerRepository {
	return &buyerRepository{
		db: db,
	}
}
func (r *buyerRepository) ClearProfilePhoto(userID uint) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"profile_image": "",
		}).Error
}

// cretae update reguser for upload buyer profile image
func (r *buyerRepository) UpdateRegUserForImage(userID uint, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return r.db.Model(&models.RegUser{}).Where("id = ?", userID).Updates(updates).Error
}

func (r *buyerRepository) GetBuyerByUserID(userID uint) (*models.Buyer, error) {
	var buyer models.Buyer
	err := r.db.Preload("RegUser").
		Where("user_id =?", userID).
		First(&buyer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("buyer not found")
		}
		return nil, fmt.Errorf("failed to fetch buyer: %v", err)

	}
	return &buyer, nil

}

func (r *buyerRepository) GetBuyerByID(buyerID uint) (*models.Buyer, error) {
	var buyer models.Buyer
	err := r.db.Where("user_id = ?", buyerID).First(&buyer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("buyer not found")
		}
		return nil, fmt.Errorf("failed to fetch buyer: %v", err)
	}
	return &buyer, nil
}
func (r *buyerRepository) GetBuyerWithRelations(buyerID uint) (*models.Buyer, error) {
	var buyer models.Buyer
	err := r.db.
		Preload("RegUser").
		Preload("Role").
		Preload("DefaultAddr").
		Preload("Addresses", func(db *gorm.DB) *gorm.DB {
			return db.Order("is_default DESC, created_at DESC")
		}).
		Where("id = ?", buyerID).
		First(&buyer).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("buyer not found")
		}
		return nil, fmt.Errorf("failed to fetch buyer: %v", err)
	}
	return &buyer, nil
}
func (r *buyerRepository) UpdateBuyerProfile(buyer *models.Buyer) error {
	tx := r.db.Begin()

	// --- 1️⃣ Buyer table updates ---
	buyerUpdate := map[string]interface{}{}

	if buyer.Phone != "" {
		buyerUpdate["phone"] = buyer.Phone
	}
	if buyer.ProfilePicture != "" {
		buyerUpdate["profile_picture"] = buyer.ProfilePicture
	}

	if len(buyerUpdate) > 0 {
		if err := tx.Model(&models.Buyer{}).
			Where("id = ?", buyer.ID).
			Updates(buyerUpdate).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// --- 2️⃣ RegUser table updates ---
	if buyer.RegUser.ID != 0 {
		userUpdate := map[string]interface{}{}

		if buyer.RegUser.FirstName != "" {
			userUpdate["first_name"] = buyer.RegUser.FirstName
		}
		if buyer.RegUser.LastName != "" {
			userUpdate["last_name"] = buyer.RegUser.LastName
		}
		if buyer.RegUser.Email != "" {
			userUpdate["email"] = buyer.RegUser.Email
		}

		if len(userUpdate) > 0 {
			if err := tx.Model(&models.RegUser{}).
				Where("id = ?", buyer.RegUser.ID).
				Updates(userUpdate).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func (r *buyerRepository) GetBuyerStats(buyerID uint) (*buyerprofile.BuyerStateResponse, error) {
	stats := &buyerprofile.BuyerStateResponse{}
	// Get basic buyer info
	var buyer models.Buyer
	if err := r.db.Where("id = ?", buyerID).First(&buyer).Error; err != nil {
		return nil, err
	}

	stats.TotalOrders = buyer.TotalOrdersCount
	stats.TotalSpent = buyer.TotalSpent
	stats.LastOrderDate = buyer.LastOrderAt

	// Count orders by status - using int64 variables for GORM Count
	var pendingOrders, completedOrders, cancelledOrders, wishlistCount, cartItemCount, reviewCount int64

	r.db.Model(&models.Order{}).
		Where("buyer_id = ? AND status = ?", buyerID, "pending").
		Count(&pendingOrders)
	stats.PendingOrders = int(pendingOrders)

	r.db.Model(&models.Order{}).
		Where("buyer_id = ? AND status = ?", buyerID, "completed").
		Count(&completedOrders)
	stats.CompletedOrders = int(completedOrders)

	r.db.Model(&models.Order{}).
		Where("buyer_id = ? AND status IN ?", buyerID, []string{"cancelled", "refunded"}).
		Count(&cancelledOrders)
	stats.CancelledOrders = int(cancelledOrders)

	// Count wishlist items
	r.db.Model(&models.Wishlist{}).
		Where("buyer_id = ?", buyerID).
		Count(&wishlistCount)
	stats.WishlistCount = int(wishlistCount)

	// Count cart items
	r.db.Model(&models.CartItem{}).
		Where("buyer_id = ?", buyerID).
		Count(&cartItemCount)
	stats.CartItemCount = int(cartItemCount)

	// Count reviews
	r.db.Model(&models.Review{}).
		Where("buyer_id = ?", buyerID).
		Count(&reviewCount)
	stats.ReviewCount = int(reviewCount)

	// Get recent orders
	var recentOrders []models.Order
	r.db.Where("buyer_id = ?", buyerID).
		Order("created_at DESC").
		Limit(5).
		Find(&recentOrders)

	for _, order := range recentOrders {
		var itemCount int64
		r.db.Model(&models.OrderItem{}).Where("order_id = ?", order.ID).Count(&itemCount)

		stats.RecentOrders = append(stats.RecentOrders, buyerprofile.RecentOrderInfo{
			OrderID:     order.ID,
			OrderNumber: order.OrderNumber,
			TotalAmount: order.Total,
			Status:      order.Status,
			ItemCount:   int(itemCount),
			OrderDate:   order.CreatedAt,
		})
	}

	return stats, nil
}

func (r *buyerRepository) GetBuyerOrders(buyerID uint, page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit

	// Count total
	r.db.Model(&models.Order{}).Where("buyer_id = ?", buyerID).Count(&total)

	// Get orders with items
	err := r.db.
		Preload("OrderItems.Product").
		Where("buyer_id = ?", buyerID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&orders).Error

	return orders, total, err
}

func (r *buyerRepository) UpdateRegUser(regUser *models.RegUser) error {
	return r.db.Save(regUser).Error
}

// Address management methods
func (r *buyerRepository) CreateAddress(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *buyerRepository) UpdateAddress(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *buyerRepository) DeleteAddress(addressID, buyerID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var address models.Address

		// Fetch the address in one query
		if err := tx.First(&address, "id = ? AND buyer_id = ?", addressID, buyerID).Error; err != nil {
			return err
		}

		// If it's default, unset default_address_id for buyer
		if address.IsDefault {
			if err := tx.Model(&models.Buyer{}).
				Where("id = ? AND default_address_id = ?", buyerID, addressID).
				Update("default_address_id", nil).Error; err != nil {
				return err
			}
		}

		// Delete the address
		if err := tx.Delete(&models.Address{}, "id = ? AND buyer_id = ?", addressID, buyerID).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *buyerRepository) GetAddressByID(addressID, buyerID uint) (*models.Address, error) {
	var address models.Address
	err := r.db.Where("id = ? AND buyer_id = ?", addressID, buyerID).First(&address).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("address not found")
		}
		return nil, err
	}
	return &address, nil
}

func (r *buyerRepository) GetBuyerAddresses(buyerID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Where("buyer_id = ?", buyerID).Order("is_default DESC, created_at DESC").Find(&addresses).Error
	return addresses, err
}

func (r *buyerRepository) SetDefaultAddress(buyerID, addressID uint) error {
	tx := r.db.Begin()

	// Unset all default addresses
	if err := tx.Model(&models.Address{}).Where("buyer_id = ?", buyerID).Update("is_default", false).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set new default
	if err := tx.Model(&models.Address{}).Where("id = ? AND buyer_id = ?", addressID, buyerID).Update("is_default", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update buyer's default address ID
	if err := tx.Model(&models.Buyer{}).Where("id = ?", buyerID).Update("default_address_id", addressID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
