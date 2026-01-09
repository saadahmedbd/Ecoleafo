package orderservice

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) UpdateOrderItemStatus(userID uint, userRole, username string, req order.UpdateOrderItemStatusRequest) error {
	item, err := s.orderrepo.GetOrderItemByID(req.ItemID)
	if err != nil {
		return err
	}

	// Only seller of the item or admin can update
	if userRole != "[seller]" && item.SellerID != userID {
		return errors.New("unauthorized")
	}

	item.Status = req.Status

	if req.Status == string(models.OrderShipped) && item.ShippedAt == nil {
		now := time.Now()
		item.ShippedAt = &now
	}

	if req.Status == string(models.OrderDelivered) && item.DeliveredAt == nil {
		now := time.Now()
		item.DeliveredAt = &now
	}

	if err := s.orderrepo.UpdateOrderItem(item); err != nil {
		return err
	}

	// Create history
	comment := req.Comment
	if comment == "" {
		comment = fmt.Sprintf("Item %s status updated to %s", item.ProductName, req.Status)
	}

	history := &models.OrderHistory{
		OrderID:   item.OrderID,
		Status:    models.OrderStatusEnum(req.Status),
		Comment:   comment,
		UpdatedBy: username,
	}
	s.orderrepo.CreateOrderHistory(history)

	return nil
}
