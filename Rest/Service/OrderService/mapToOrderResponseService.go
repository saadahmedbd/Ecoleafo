package orderservice

import (
	"math"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) mapToOrderResponse(orderModel *models.Order, userRole string) (*order.OrderResponse, error) {

	// Prepare Order Item Responses
	items := make([]order.OrderItemResponse, len(orderModel.OrderItems))

	for i, item := range orderModel.OrderItems {

		// ---- IMAGE LOGIC ----
		var imageURL string

		// Primary Image
		for _, img := range item.Product.Images {
			if img.IsPrimary {
				imageURL = img.ImageURL
				break
			}
		}

		// If no primary found → use first image
		if imageURL == "" && len(item.Product.Images) > 0 {
			imageURL = item.Product.Images[0].ImageURL
		}

		// ---- SELLER NAME ----
		sellerName := ""
		if item.Seller.ID != 0 {
			sellerName = item.Seller.StoreName
		}

		// ---- COMMISSION LOGIC ----
		var commission float64
		var sellerEarning float64

		if userRole == "admin" || userRole == "seller" {
			commission = item.Commission
			sellerEarning = item.SellerEarning
		}

		// ---- BUILD RESPONSE ----
		items[i] = order.OrderItemResponse{
			ID:            item.ID,
			ProductID:     item.ProductID,
			ProductName:   item.ProductName,
			ProductSKU:    item.ProductSKU,
			SellerID:      item.SellerID,
			SellerName:    sellerName,
			Quantity:      item.Quantity,
			Price:         math.Round(item.Price),
			Total:         math.Round(item.Total),
			DiscountPrice: math.Round(item.Product.DiscountPrice),
			Image:         imageURL,
			Commission:    commission,
			SellerEarning: sellerEarning,
			Status:        item.Status,
			ShippedAt:     item.ShippedAt,
			DeliveredAt:   item.DeliveredAt,
		}
	}

	// ---- ORDER RESPONSE ----
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
		OrderItems:      items,
	}

	return resp, nil
}
