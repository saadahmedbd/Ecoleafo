package orderservice

import order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"

func (s *orderService) GetAllOrders(filter order.OrderListFilter) ([]order.OrderResponse, int64, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	filterMap := make(map[string]interface{})
	if filter.Status != "" {
		filterMap["status"] = filter.Status
	}
	if filter.PaymentStatus != "" {
		filterMap["payment_status"] = filter.PaymentStatus
	}
	if filter.BuyerID > 0 {
		filterMap["buyer_id"] = filter.BuyerID
	}

	orders, total, err := s.orderrepo.FindAll(filterMap, filter.Page, filter.Limit)
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
