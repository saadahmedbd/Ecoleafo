package orderservice

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) CancelOrder(orderID, userID uint, userRole, username string) (*order.OrderResponse, error) {
	order, err := s.orderrepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	// Check authorization
	if userRole != "[buyer]" && order.BuyerID != userID {
		return nil, errors.New("unauthorized")
	}

	// Can only cancel pending or processing orders
	if order.Status != string(models.OrderPending) && order.Status != string(models.OrderProcessing) {
		return nil, errors.New("cannot cancel order in current status")
	}

	order.Status = string(models.OrderCancelled)

	if err := s.orderrepo.Update(order); err != nil {
		return nil, err
	}

	// Restore product stock
	items, _ := s.orderrepo.GetOrderItemByOrderID(orderID)
	for _, item := range items {
		product, _ := s.productservice.GetProduct(item.ProductID)
		if product != nil {
			product.Quantity += item.Quantity
			s.orderrepo.UpdateStock(product.ID, product.Quantity)
		}

		// Update item status
		item.Status = string(models.OrderCancelled)
		s.orderrepo.UpdateOrderItem(&item)
	}

	// Create history
	history := &models.OrderHistory{
		OrderID:   order.ID,
		Status:    models.OrderCancelled,
		Comment:   "Order cancelled",
		UpdatedBy: username,
	}
	s.orderrepo.CreateOrderHistory(history)

	return s.mapToOrderResponse(order)
}
