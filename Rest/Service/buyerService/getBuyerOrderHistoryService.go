package buyerservice

import (
	"math"

	buyerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/BuyerProfile"
)

func (s *buyerService) GetBuyerOrderHistory(userIDFromJWT uint, page, limit int) (*buyerprofile.BuyerOrderHistoryResponse, error) {
	buyer, err := s.buyerRepo.GetBuyerByUserID(userIDFromJWT)
	if err != nil {
		return nil, err
	}

	orders, total, err := s.buyerRepo.GetBuyerOrders(buyer.ID, page, limit)
	if err != nil {
		return nil, err
	}

	var orderInfos []buyerprofile.OrderInfo
	for _, order := range orders {
		var items []buyerprofile.OrderItemInfo
		for _, item := range order.OrderItems {
			items = append(items, buyerprofile.OrderItemInfo{
				ProductID:   item.ProductID,
				ProductName: item.Product.Name,
				Quantity:    item.Quantity,
				Price:       item.Price,
				Subtotal:    item.Order.Subtotal,
			})
		}

		orderInfos = append(orderInfos, buyerprofile.OrderInfo{
			ID:          order.ID,
			OrderNumber: order.OrderNumber,
			TotalAmount: order.Total,
			Status:      order.Status,
			ItemCount:   len(order.OrderItems),
			OrderDate:   order.CreatedAt,
			Items:       items,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &buyerprofile.BuyerOrderHistoryResponse{
		Orders: orderInfos,
		Pagination: buyerprofile.PaginationInfo{
			Page:       page,
			Limit:      limit,
			Total:      int(total),
			TotalPages: totalPages,
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	}, nil
}
