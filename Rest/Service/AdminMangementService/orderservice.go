package adminmangementservice

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	adminmangement "github.com/saadahmedbd/Treestore/Rest/Repository/AdminMangement"
	audithelper "github.com/saadahmedbd/Treestore/Rest/Service/AuditHelper"
	"github.com/saadahmedbd/Treestore/constants"
)

type OrderService struct {
	orderRepo       *adminmangement.OrderRepository
	activityLogRepo *adminmangement.AuditLogRepository
	auditHelper     *audithelper.AuditHelper
}

func NewOrderService(orderRepo *adminmangement.OrderRepository,
	activityLogRepo *adminmangement.AuditLogRepository,
	auditHelper *audithelper.AuditHelper,
) *OrderService {
	return &OrderService{
		orderRepo:       orderRepo,
		activityLogRepo: activityLogRepo,
		auditHelper:     auditHelper,
	}
}

// GetAllOrders retrieves all orders with pagination
func (s *OrderService) GetAllOrders(page, limit int, status string) ([]models.Order, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.orderRepo.GetAll(page, limit, status)
}

// GetOrderByID retrieves an order by ID
func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.orderRepo.GetByID(id)
}

// GetOrderByOrderNumber retrieves an order by order number
func (s *OrderService) GetOrderByOrderNumber(orderNumber string) (*models.Order, error) {
	return s.orderRepo.GetByOrderNumber(orderNumber)
}

// UpdateOrderStatus updates order status
func (s *OrderService) UpdateOrderStatus(id uint, status string, adminID uint) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	if order.Status == status {
		return errors.New("order is already in this status")
	}

	// Validate status transition
	validStatuses := []string{
		constants.OrderStatusPending,
		constants.OrderStatusConfirmed,
		constants.OrderStatusProcessing,
		constants.OrderStatusShipped,
		constants.OrderStatusDelivered,
		constants.OrderStatusCancelled,
		constants.OrderStatusRefunded,
	}

	isValid := false
	for _, s := range validStatuses {
		if s == status {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("invalid order status")
	}

	err = s.orderRepo.UpdateStatus(id, status)
	if err != nil {
		return err
	}
	//log order status

	// Log activity
	s.logOrderActivity(adminID, constants.ActionUpdate, &id, "Updated order status to "+status)

	return nil
}

// SearchOrders searches orders by query
func (s *OrderService) SearchOrders(query string, page, limit int) ([]models.Order, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.orderRepo.Search(query, page, limit)
}

// CancelOrder cancels an order
func (s *OrderService) CancelOrder(id uint, reason string, adminID uint) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	if order.Status == constants.OrderStatusCancelled {
		return errors.New("order is already cancelled")
	}

	if order.Status == constants.OrderStatusDelivered {
		return errors.New("delivered orders cannot be cancelled")
	}

	order.Status = constants.OrderStatusCancelled
	order.CancellationReason = reason

	err = s.orderRepo.Update(order)
	if err != nil {
		return err
	}
	//log order status

	// Log activity
	s.logOrderActivity(adminID, constants.ActionUpdate, &id, "Cancelled order: "+reason)

	return nil
}

// GetOrderStats retrieves order statistics
func (s *OrderService) GetOrderStats() (map[string]interface{}, error) {
	return s.orderRepo.GetStats()
}

// GetRecentOrders retrieves recent orders
func (s *OrderService) GetRecentOrders(limit int) ([]models.Order, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}
	return s.orderRepo.GetRecentOrders(limit)
}

// Helper function
func (s *OrderService) logOrderActivity(adminID uint, action string, entityID *uint, description string) {
	log := &models.AuditLog{
		ActorID:     &adminID,
		Action:      action,
		EntityType:  "order",
		EntityID:    entityID,
		Description: description,
	}
	s.activityLogRepo.Create(log)
}
