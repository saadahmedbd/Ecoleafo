package orderservice

import (
	"errors"
	"fmt"
	"math"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateOrder creates a new order with proper transaction handling and security
func (s *orderService) CreateOrder(buyerID uint, req order.CreateOrderRequest) (*order.OrderResponse, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	buyer, err := s.buyerProfilerepo.GetBuyerByID(buyerID)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("buyer not found")
	}

	cartItems, err := s.cartitemrepo.GetSelectedCartItems(buyerID)
	if err != nil || len(cartItems) == 0 {
		tx.Rollback()
		return nil, errors.New("no items selected for checkout")
	}

	productIDs := extractProductIDs(cartItems)
	products, err := s.validateAndLockInventory(tx, productIDs, cartItems)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	pricing, err := s.calculateOrderPricing(buyerID, cartItems, products, req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if pricing.Total <= 0 {
		tx.Rollback()
		return nil, errors.New("invalid order total")
	}

	orderNumber := s.generateOrderNumber(buyerID)

	newOrder := models.Order{
		OrderNumber:     orderNumber,
		BuyerID:         buyerID,
		Status:          string(models.OrderPending),
		PaymentStatus:   "pending",
		PaymentMethod:   req.PaymentMethod,
		Subtotal:        pricing.Subtotal,
		ShippingCost:    pricing.ShippingCost,
		DiscountAmount:  pricing.DiscountAmount,
		TaxAmount:       pricing.TaxAmount,
		Total:           pricing.Total,
		ShippingAddress: req.ShippingAddress,
		BillingAddress:  req.BuillingAddress,
		CustomerEmail:   req.CustomerEmail,
		CustomerPhone:   req.CustomerPhone,
		Notes:           req.Notes,
	}

	if err := tx.Create(&newOrder).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	if err := s.createOrderItems(tx, &newOrder, cartItems, products); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := s.updateProductStock(tx, cartItems, products); err != nil {
		tx.Rollback()
		return nil, err
	}

	history := &models.OrderHistory{
		OrderID:   newOrder.ID,
		Status:    models.OrderPending,
		Comment:   "Order created successfully",
		UpdatedBy: fmt.Sprintf("buyer_%d", buyerID),
	}
	if err := tx.Create(history).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Remove only selected items from cart
	for _, item := range cartItems {
		if err := tx.Where("buyer_id = ? AND product_id = ?", buyerID, item.ProductID).Delete(&models.CartItem{}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	go s.sendOrderConfirmationEmail(&newOrder, buyer)
	go s.notifySellers(&newOrder)
	go s.recordOrderMetrics(&newOrder)

	return s.mapToOrderResponse(&newOrder, "buyer")
}

func (s *orderService) validateAndLockInventory(tx *gorm.DB, productIDs []uint, cartItems []models.CartItem) (map[uint]*models.Product, error) {
	var products []models.Product
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to lock products: %w", err)
	}

	productMap := make(map[uint]*models.Product)
	for i := range products {
		productMap[products[i].ID] = &products[i]
	}

	for _, item := range cartItems {
		product, exists := productMap[item.ProductID]
		if !exists {
			return nil, fmt.Errorf("product %d not found", item.ProductID)
		}
		if !product.IsActive {
			return nil, fmt.Errorf("product '%s' is no longer available", product.Name)
		}
		if product.Quantity < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for '%s' (available: %d, requested: %d)", product.Name, product.Quantity, item.Quantity)
		}
	}

	return productMap, nil
}

type PricingDetails struct {
	Subtotal       float64
	ShippingCost   float64
	DiscountAmount float64
	Total          float64
	TaxAmount      float64
}

// Industry Standard Calculation Order:
// 1. Calculate item subtotals (price × quantity)
// 2. Apply product-level discounts
// 3. Calculate subtotal after discounts
// 4. Calculate shipping (based on discounted subtotal)
// 5. Calculate tax (on discounted subtotal + shipping)
// 6. Calculate final total
func (s *orderService) calculateOrderPricing(buyerID uint, cartItems []models.CartItem, products map[uint]*models.Product, req order.CreateOrderRequest) (*PricingDetails, error) {
	var subtotal float64
	var totalDiscountAmount float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		originalPrice := product.Price
		itemSubtotal := originalPrice * float64(item.Quantity)
		subtotal += itemSubtotal

		if product.DiscountPrice > 0 {
			itemDiscount := (originalPrice - product.DiscountPrice) * float64(item.Quantity)
			totalDiscountAmount += itemDiscount
		}
	}

	shippingCost, err := s.calculateShippingCost(req.ShippingAddress, req.ShippingMethodID, cartItems, products)
	if err != nil {
		return nil, fmt.Errorf("shipping calculation failed: %w", err)
	}

	total := subtotal - totalDiscountAmount + shippingCost

	return &PricingDetails{
		Subtotal:       math.Round(subtotal),
		DiscountAmount: math.Round(totalDiscountAmount),
		ShippingCost:   shippingCost,
		Total:          math.Round(total),
	}, nil
}

func (s *orderService) calculateShippingCost(address string, methodID uint, cartItems []models.CartItem, products map[uint]*models.Product) (float64, error) {
	var subtotal float64
	for _, item := range cartItems {
		product := products[item.ProductID]
		subtotal += product.Price * float64(item.Quantity)
	}

	if subtotal >= s.config.FreeShippingThreshold {
		return 0, nil
	}

	baseCost := 100.0
	zone := s.getShippingZone(address)
	cost := baseCost * s.getZoneMultiplier(zone)

	return cost, nil
}

func (s *orderService) createOrderItems(tx *gorm.DB, order *models.Order, cartItems []models.CartItem, products map[uint]*models.Product) error {
	commissionRate := s.config.CommissionRate

	for _, item := range cartItems {
		product := products[item.ProductID]
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

		if err := tx.Create(orderItem).Error; err != nil {
			return fmt.Errorf("failed to create order item: %w", err)
		}
	}

	return nil
}

func (s *orderService) updateProductStock(tx *gorm.DB, cartItems []models.CartItem, products map[uint]*models.Product) error {
	for _, item := range cartItems {
		product := products[item.ProductID]
		newQuantity := product.Quantity - item.Quantity

		if err := tx.Model(&models.Product{}).Where("id = ?", product.ID).Update("quantity", newQuantity).Error; err != nil {
			return fmt.Errorf("failed to update stock for product %d: %w", product.ID, err)
		}
	}

	return nil
}

func (s *orderService) markCouponAsUsed(tx *gorm.DB, couponCode string, buyerID, orderID uint) error {
	return nil
}

func (s *orderService) generateOrderNumber(buyerID uint) string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("ORD-%d-%d", timestamp, buyerID)
}

func extractProductIDs(cartItems []models.CartItem) []uint {
	ids := make([]uint, len(cartItems))
	for i, item := range cartItems {
		ids[i] = item.ProductID
	}
	return ids
}

func (s *orderService) sendOrderConfirmationEmail(order *models.Order, buyer *models.Buyer) {
}

func (s *orderService) notifySellers(order *models.Order) {
}

func (s *orderService) recordOrderMetrics(order *models.Order) {
}
