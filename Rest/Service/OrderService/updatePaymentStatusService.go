package orderservice

import (
	"errors"
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) UpdatePaymentStatus(orderID, userID uint, userRole, username string, req order.UpdatePaymentStatusRequest) (*order.OrderResponse, error) {
	order, err := s.orderrepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	// Only admins and sellers can update payment status
	if userRole == "[buyer]" {
		return nil, errors.New("unauthorized")
	}

	order.PaymentStatus = req.PaymentStatus

	if err := s.orderrepo.Update(order); err != nil {
		return nil, err
	}

	// Create history
	history := &models.OrderHistory{
		OrderID:   order.ID,
		Status:    models.OrderStatusEnum(order.Status),
		Comment:   fmt.Sprintf("Payment status updated to %s", req.PaymentStatus),
		UpdatedBy: username,
	}
	s.orderrepo.CreateOrderHistory(history)

	return s.mapToOrderResponse(order, userRole)
}
