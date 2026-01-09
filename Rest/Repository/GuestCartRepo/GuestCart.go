package guestcartrepo

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type GuestCartRepository interface {
	AddGuestCartItem(item *models.GuestCartItem) error
	GetGuestCart(sessionID string) ([]models.GuestCartItem, error)
	UpdateGuestCartItem(item *models.GuestCartItem) error
	RemoveGuestCartItem(sessionID string, productID uint) error
	ClearGuestCart(sessionID string) error
	MigrateGuestCartToUser(sessionID string, buyerID uint) error

	//registered user cart opreartion
	AddCartItem(item *models.CartItem) error
	GetUserCart(buyerID uint) ([]models.CartItem, error)
	UpdateCartItem(item *models.CartItem) error
	RemoveCartItem(buyerID uint, productID uint) error
	ClearUserCart(buyerID uint) error
}

type guestCartRepository struct {
	db *gorm.DB
}

func NewGuestCartRepository(db *gorm.DB) GuestCartRepository {
	return &guestCartRepository{
		db: db,
	}
}

// guest cart method
func (r *guestCartRepository) AddGuestCartItem(item *models.GuestCartItem) error {
	// check item already exist in guest cart item
	var existingItem models.GuestCartItem
	err := r.db.Where("session_id = ? AND product_id = ?",
		item.SessionID, item.ProductID).
		First(&existingItem).Error
	if err == nil {
		// item exists, update quantity
		existingItem.Quantity += item.Quantity
		return r.db.Save(&existingItem).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// item does not exist, create new
		return r.db.Create(item).Error
	}
	return err
}
func (r *guestCartRepository) GetGuestCart(sessionID string) ([]models.GuestCartItem, error) {
	var items []models.GuestCartItem
	err := r.db.Preload("Product").
		Preload("Product.Images", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_primary = ?", true).Limit(1)
		}).
		Where("session_id = ?", sessionID).
		Find(&items).Error
	return items, err
}
func (r *guestCartRepository) UpdateGuestCartItem(item *models.GuestCartItem) error {
	return r.db.Save(item).Error
}
func (r *guestCartRepository) RemoveGuestCartItem(sessionID string, productID uint) error {
	return r.db.Where("session_id= ? AND product_id = ?", sessionID, productID).
		Delete(&models.GuestCartItem{}).Error
}

func (r *guestCartRepository) ClearGuestCart(sessionID string) error {
	return r.db.Where("session_id = ?", sessionID).
		Delete(&models.GuestCartItem{}).Error
}
func (r *guestCartRepository) MigrateGuestCartToUser(sessionID string, buyerID uint) error {
	var guestItems []models.GuestCartItem
	err := r.db.Where("session_id = ?", sessionID).Find(&guestItems).Error
	if err != nil {
		return err
	}
	//migarte to user cart
	for _, guestItem := range guestItems {
		userItem := models.CartItem{
			BuyerID:   buyerID,
			ProductID: guestItem.ProductID,
			Quantity:  guestItem.Quantity,
			Price:     guestItem.Price,
		}
		//check if the item already exist in user cart
		var existingItem models.CartItem
		err := r.db.Where("buyer_id = ? AND product_id = ?", buyerID, guestItem.ProductID).
			First(&existingItem).Error
		if err == nil {
			//item exist update quantity
			existingItem.Quantity += guestItem.Quantity
			r.db.Save(&existingItem)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			//item not exist create new
			r.db.Create(&userItem)
		}
	}
	return r.ClearGuestCart(sessionID)

}

// user cart method
func (r *guestCartRepository) AddCartItem(item *models.CartItem) error {
	// check item already exist in user cart item
	var existingItem models.CartItem
	err := r.db.Where("buyer_id = ? AND product_id = ?",
		item.BuyerID, item.ProductID).
		First(&existingItem).Error
	if err == nil {
		// item exists, update quantity
		existingItem.Quantity += item.Quantity
		return r.db.Save(&existingItem).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// item does not exist, create new
		return r.db.Create(item).Error
	}
	return err
}

func (r *guestCartRepository) GetUserCart(buyerID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.db.Preload("Product").
		Preload("Product.Images", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_primary = ?", true).Limit(1)
		}).
		Where("buyer_id = ?", buyerID).
		Find(&items).Error
	return items, err
}
func (r *guestCartRepository) UpdateCartItem(item *models.CartItem) error {
	return r.db.Save(item).Error
}

func (r *guestCartRepository) RemoveCartItem(buyerID uint, productID uint) error {
	return r.db.Where("buyer_id= ? AND product_id = ?", buyerID, productID).
		Delete(&models.CartItem{}).Error
}
func (r *guestCartRepository) ClearUserCart(buyerID uint) error {
	return r.db.Where("buyer_id = ?", buyerID).
		Delete(&models.CartItem{}).Error
}
