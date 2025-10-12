package orderservice

import (
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	buyerProfilerepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerProfileRepo"
	cartitemrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CartItemRepo"
	orderrepo "github.com/saadahmedbd/Treestore/Rest/Repository/OrderRepo"
	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
)

type OrderService interface {
	CreateOrder(userID uint, req order.CreateOrderRequest) (*order.OrderResponse, error)
	GetOrderByID(orderID, userID uint, userRole string) (*order.OrderResponse, error)
	GetOrderByOrderNumber(orderNumber string, userID uint, userRole string) (*order.OrderResponse, error)
	GetBuyerOrders(buyerID uint, page, limit int) ([]order.OrderResponse, int64, error)
	GetSellerOrders(sellerID uint, page, limit int) ([]order.OrderResponse, int64, error)
	UpdateOrderStatus(orderID, userID uint, userRole, username string, req order.UpdateOrderStatusRequest) (*order.OrderResponse, error)
	UpdatePaymentStatus(orderID, userID uint, userRole, username string, req order.UpdatePaymentStatusRequest) (*order.OrderResponse, error)
	UpdateOrderItemStatus(userID uint, userRoleole, username string, req order.UpdateOrderItemStatusRequest) error
	CancelOrder(orderID, userID uint, userRole, username string) (*order.OrderResponse, error)
	GetAllOrders(filter order.OrderListFilter) ([]order.OrderResponse, int64, error)
	GetOrderHistory(orderID, userID uint, userRole string) ([]order.OrderHistoryResponse, error)
}
type orderService struct {
	orderrepo        orderrepo.OrderRepository
	cartitemrepo     cartitemrepo.CartRepository
	buyerProfilerepo buyerProfilerepo.BuyerRepository
	productservice   productservice.ProductService
}

func NewOrderService(
	OrderRepo orderrepo.OrderRepository,
	CartRepo cartitemrepo.CartRepository,
	BuyerRepo buyerProfilerepo.BuyerRepository,
	ProductService productservice.ProductService,
) OrderService {
	return &orderService{
		orderrepo:        OrderRepo,
		cartitemrepo:     CartRepo,
		buyerProfilerepo: BuyerRepo,
		productservice:   ProductService,
	}
}
