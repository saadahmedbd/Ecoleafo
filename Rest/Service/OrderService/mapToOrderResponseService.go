package orderservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) mapToOrderResponse(orderModel *models.Order) (*order.OrderResponse, error) {
	resp := &order.OrderResponse{
		ID:              orderModel.ID,
		OrderNumber:     orderModel.OrderNumber,
		BuyerID:         orderModel.BuyerID,
		Status:          orderModel.Status,
		PaymentStatus:   orderModel.PaymentStatus,
		PaymentMethod:   orderModel.PaymentMethod,
		Subtotal:        orderModel.Subtotal,
		ShippingCost:    orderModel.ShippingCost,
		TaxAmount:       orderModel.TaxAmount,
		DiscountAmount:  orderModel.DiscountAmount,
		Total:           orderModel.Total,
		ShippingAddress: orderModel.ShippingAddress,
		BillingAddress:  orderModel.BillingAddress,
		CustomerEmail:   orderModel.CustomerEmail,
		CustomerPhone:   orderModel.CustomerPhone,
		TrackingNumber:  orderModel.TrackingNumber,
		ShippedAt:       orderModel.ShippedAt,
		DeliveredAt:     orderModel.DeliveredAt,
		Notes:           orderModel.Notes,
		CreatedAt:       orderModel.CreatedAt,
		UpdatedAt:       orderModel.UpdatedAt,
	}

	if len(orderModel.OrderItems) > 0 {
		items := make([]order.OrderItemResponse, len(orderModel.OrderItems))
		for i, item := range orderModel.OrderItems {
			sellerName := ""
			if item.Seller.ID != 0 {
				sellerName = item.Seller.StoreName
			}

			items[i] = order.OrderItemResponse{
				ID:            item.ID,
				ProductID:     item.ProductID,
				ProductName:   item.ProductName,
				ProductSKU:    item.ProductSKU,
				SellerID:      item.SellerID,
				SellerName:    sellerName,
				Quantity:      item.Quantity,
				Price:         item.Price,
				Total:         item.Total,
				Commission:    item.Commission,
				SellerEarning: item.SellerEarning,
				Status:        item.Status,
				ShippedAt:     item.ShippedAt,
				DeliveredAt:   item.DeliveredAt,
			}
		}
		resp.OrderItems = items
	}

	return resp, nil
}
