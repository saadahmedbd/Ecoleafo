package orderservice

import (
	"errors"

	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) GetOrderByID(orderID, userID uint, userRole string) (*order.OrderResponse, error) {
	order, err := s.orderrepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}
	// Authorization check
	if userRole != "[buyer]" && order.BuyerID != userID {
		return nil, errors.New("unauthorized access")
	}
	if userRole == "[seller]" {
		hasAccess := false
		for _, item := range order.OrderItems {
			if item.SellerID == userID {
				hasAccess = true
				break
			}
		}
		if !hasAccess {
			return nil, errors.New("unauthorized access")
		}
	}
	return s.mapToOrderResponse(order, userRole)
}
