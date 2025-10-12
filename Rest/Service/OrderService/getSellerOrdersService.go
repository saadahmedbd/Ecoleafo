package orderservice

import order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"

func (s *orderService) GetSellerOrders(sellerID uint, page, limit int) ([]order.OrderResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	orders, total, err := s.orderrepo.FindBySellerID(sellerID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]order.OrderResponse, len(orders))
	for i, order := range orders {
		resp, _ := s.mapToOrderResponse(&order)
		responses[i] = *resp
	}

	return responses, total, nil
}
