package orderservice

import order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"

func (s *orderService) GetOrderHistory(orderID, userID uint, userRole string) ([]order.OrderHistoryResponse, error) {
	// Verify access to order
	_, err := s.GetOrderByID(orderID, userID, userRole)
	if err != nil {
		return nil, err
	}

	history, err := s.orderrepo.GetOrderHistory(orderID)
	if err != nil {
		return nil, err
	}

	responses := make([]order.OrderHistoryResponse, len(history))
	for i, h := range history {
		responses[i] = order.OrderHistoryResponse{
			ID:        h.ID,
			Status:    string(h.Status),
			Comment:   h.Comment,
			UpdatedBy: h.UpdatedBy,
			CreatedAt: h.CreatedAt,
		}
	}

	return responses, nil
}
