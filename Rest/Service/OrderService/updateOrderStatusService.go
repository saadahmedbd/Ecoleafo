package orderservice

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) UpdateOrderStatus(orderID, userID uint, userRole, username string, req order.UpdateOrderStatusRequest) (*order.OrderResponse, error) {
	order, err := s.orderrepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	// Only sellers and admins can update order status
	if userRole == "[buyer]" {
		return nil, errors.New("buyers cannot update order status")
	}

	oldStatus := order.Status
	order.Status = req.Status

	if req.TrackingNumber != "" {
		order.TrackingNumber = req.TrackingNumber
	}

	if req.Status == string(models.OrderShipped) && order.ShippedAt == nil {
		now := time.Now()
		order.ShippedAt = &now
	}

	if req.Status == string(models.OrderDelivered) && order.DeliveredAt == nil {
		now := time.Now()
		order.DeliveredAt = &now
	}

	if err := s.orderrepo.Update(order); err != nil {
		return nil, err
	}

	// Update all order items status
	items, _ := s.orderrepo.GetOrderItemByOrderID(orderID)
	for _, item := range items {
		item.Status = req.Status
		if req.Status == string(models.OrderShipped) && item.ShippedAt == nil {
			now := time.Now()
			item.ShippedAt = &now
		}
		if req.Status == string(models.OrderDelivered) && item.DeliveredAt == nil {
			now := time.Now()
			item.DeliveredAt = &now
		}
		s.orderrepo.UpdateOrderItem(&item)
	}

	// Create history
	comment := req.Comment
	if comment == "" {
		comment = fmt.Sprintf("Status changed from %s to %s", oldStatus, req.Status)
	}

	history := &models.OrderHistory{
		OrderID:   order.ID,
		Status:    models.OrderStatusEnum(req.Status),
		Comment:   comment,
		UpdatedBy: username,
	}
	s.orderrepo.CreateOrderHistory(history)

	return s.mapToOrderResponse(order, userRole)
}
