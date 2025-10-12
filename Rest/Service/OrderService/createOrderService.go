package orderservice

import (
	"errors"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
)

func (s *orderService) CreateOrder(buyerID uint, req order.CreateOrderRequest) (*order.OrderResponse, error) {
	// get buyer info
	_, err := s.buyerProfilerepo.GetBuyerByID(buyerID)
	if err != nil {
		return nil, errors.New("buyer not found")
	}
	//get cart items
	cartItems, err := s.cartitemrepo.GetCart(buyerID, true)
	if err != nil || len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}
	//calculate subtotal
	var subtotal float64
	for _, item := range cartItems {

		product, err := s.productservice.GetProduct(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product %d not found", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s", product.Name)
		}
		subtotal += product.Price * float64(item.Quantity)
	}
	// calculate total
	total := subtotal + req.ShippingCost - req.DiscountAmount

	//generate order number
	orderNumber := fmt.Sprintf("ORD-%d-%d", time.Now().Unix(), buyerID)
	//create order
	order := models.Order{
		OrderNumber:     orderNumber,
		BuyerID:         buyerID,
		Status:          string(models.OrderPending),
		PaymentStatus:   "pending",
		PaymentMethod:   req.PaymentMethod,
		Subtotal:        subtotal,
		ShippingCost:    req.ShippingCost,
		DiscountAmount:  req.DiscountAmount,
		Total:           total,
		ShippingAddress: req.ShippingAddress,
		BillingAddress:  req.BuillingAddress,
		CustomerEmail:   req.CustomerEmail,
		CustomerPhone:   req.CustomerPhone,
		Notes:           req.Notes,
	}
	if err := s.orderrepo.Create(&order); err != nil {
		return nil, err
	}
	//platform commission rate
	const commissionRate = 0.15
	//cretae order item and update stock
	for _, item := range cartItems {
		product, _ := s.productservice.GetProduct(item.ProductID)

		itemTotal := product.Price * float64(item.Quantity)
		commission := itemTotal * commissionRate
		sellerEarning := itemTotal - commission

		orderItem := &models.OrderItem{
			OrderID:       order.ID,
			ProductID:     item.ProductID,
			SellerID:      product.SellerID,
			ProductName:   product.Name,
			ProductSKU:    product.SKU,
			Quantity:      item.Quantity,
			Price:         product.Price,
			Total:         itemTotal,
			Commission:    commission,
			SellerEarning: sellerEarning,
			Status:        string(models.OrderPending),
		}
		if err := s.orderrepo.CreateOrderItem(orderItem); err != nil {
			return nil, err
		}
		//update product quamtity/stock
		product.Quantity -= item.Quantity
		if err := s.orderrepo.UpdateStock(product.ID, product.Quantity); err != nil {

			return nil, err
		}

	}
	//create order history
	history := &models.OrderHistory{
		OrderID:   order.ID,
		Status:    models.OrderPending,
		Comment:   "order created",
		UpdatedBy: fmt.Sprintf("buyer_%d", buyerID),
	}
	s.orderrepo.CreateOrderHistory(history)

	//clear cart
	s.cartitemrepo.ClearCart(buyerID)
	return s.mapToOrderResponse(&order)

}
