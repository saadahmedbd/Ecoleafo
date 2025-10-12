package orderrepo

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	FindByID(id uint) (*models.Order, error)
	FindByOrderNumber(orderNumber string) (*models.Order, error)
	FindByBuyerID(buyerID uint, page, limit int) ([]models.Order, int64, error)
	FindBySellerID(sellerID uint, page, limit int) ([]models.Order, int64, error)
	FindAll(filter map[string]interface{}, page, limit int) ([]models.Order, int64, error)
	Update(order *models.Order) error
	Delete(id uint) error
	CreateOrderItem(item *models.OrderItem) error
	UpdateOrderItem(item *models.OrderItem) error
	CreateOrderHistory(history *models.OrderHistory) error
	GetOrderItemByOrderID(orderID uint) ([]models.OrderItem, error)
	GetOrderItemByID(itemID uint) (*models.OrderItem, error)
	GetOrderHistory(orderID uint) ([]models.OrderHistory, error)

	UpdateStock(id uint, newStock int) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}
