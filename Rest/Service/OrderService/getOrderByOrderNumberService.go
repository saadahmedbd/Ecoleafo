package orderservice

import (
	"errors"

	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) GetOrderByOrderNumber(orderNumber string, userID uint, userRole string) (*order.OrderResponse, error) {
	order, err := s.orderrepo.FindByOrderNumber(orderNumber)
	if err != nil {
		return nil, err
	}
	//authorization check
	if userRole != "[buyer]" && order.BuyerID != userID {
		return nil, errors.New("unauthorized access")
	}
	return s.mapToOrderResponse(order)
}
