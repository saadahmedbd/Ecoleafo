package messagingrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (r *messagingRepositoryImpl) GetBuyerInfo(buyerID uint) (*models.Buyer, error) {
	var buyer models.Buyer
	err := r.db.Preload("RegUser").First(&buyer, buyerID).Error
	return &buyer, err
}
func (r *messagingRepositoryImpl) GetSellerInfo(sellerID uint) (*models.User, error) {
	var seller models.User
	err := r.db.Preload("RegUser").First(&seller, sellerID).Error
	return &seller, err
}
func (r *messagingRepositoryImpl) GetAdminInfo(adminID uint) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Preload("RegUser").First(&admin, adminID).Error
	return &admin, err
}
func (r *messagingRepositoryImpl) GetRegUserInfo(userID uint) (*models.RegUser, error) {
	var regUser models.RegUser
	err := r.db.First(&regUser, userID).Error
	return &regUser, err
}

func (r *messagingRepositoryImpl) GetProductInfo(productID uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC").Limit(1)
	}).First(&product, productID).Error
	return &product, err
}
func (r *messagingRepositoryImpl) GetOrderInfo(orderID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, orderID).Error
	return &order, err
}
