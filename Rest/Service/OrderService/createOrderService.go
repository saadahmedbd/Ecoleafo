package orderservice

import (
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	order "github.com/saadahmedbd/Treestore/Rest/DTO/Order"
	util "github.com/saadahmedbd/Treestore/Util"
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

	// Load RegUser to get phone if buyer.Phone is empty
	if err := tx.Preload("RegUser").First(&buyer, buyerID).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to load buyer details")
	}

	// Fetch phone numbers from addresses if IDs provided
	shippingPhone := req.ShippingPhoneNumber
	if shippingPhone == "" && req.ShippingAddressID != nil {
		var shippingAddr models.Address
		if err := tx.First(&shippingAddr, "id = ? AND buyer_id = ?", *req.ShippingAddressID, buyerID).Error; err == nil {
			shippingPhone = shippingAddr.Phone
		}
	}

	billingPhone := req.BillingPhoneNumber
	if billingPhone == "" && req.BillingAddressID != nil {
		var billingAddr models.Address
		if err := tx.First(&billingAddr, "id = ? AND buyer_id = ?", *req.BillingAddressID, buyerID).Error; err == nil {
			billingPhone = billingAddr.Phone
		}
	}

	// If still empty, try to get from default address
	if shippingPhone == "" || billingPhone == "" {
		var defaultAddr models.Address
		if err := tx.First(&defaultAddr, "buyer_id = ? AND is_default = ?", buyerID, true).Error; err == nil {
			if shippingPhone == "" {
				shippingPhone = defaultAddr.Phone
			}
			if billingPhone == "" {
				billingPhone = defaultAddr.Phone
			}
		}
	}

	// Fallback to buyer/regUser phone if still empty
	defaultPhone := "N/A"
	if buyer.Phone != "" && buyer.Phone != "N/A" {
		defaultPhone = buyer.Phone
	} else if buyer.RegUser != nil && buyer.RegUser.Phone != "" && buyer.RegUser.Phone != "N/A" {
		defaultPhone = buyer.RegUser.Phone
	}

	if shippingPhone == "" {
		shippingPhone = defaultPhone
	}
	if billingPhone == "" {
		billingPhone = defaultPhone
	}

	customerPhone := req.CustomerPhone
	if customerPhone == "" {
		customerPhone = defaultPhone
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

	giftCharge := 0.0
	isGift := req.IsGift
	for _, item := range cartItems {
		if item.IsGift {
			isGift = true
			break
		}
	}
	if isGift {
		giftCharge = 50.0
		pricing.Total += giftCharge
	}

	if pricing.Total <= 0 {
		tx.Rollback()
		return nil, errors.New("invalid order total")
	}

	orderNumber := s.generateOrderNumber(buyerID)

	newOrder := models.Order{
		OrderNumber:          orderNumber,
		BuyerID:              buyerID,
		Status:               string(models.OrderPending),
		PaymentStatus:        "pending",
		PaymentMethod:        req.PaymentMethod,
		Subtotal:             pricing.Subtotal,
		ShippingCost:         pricing.ShippingCost,
		DiscountAmount:       pricing.DiscountAmount,
		TaxAmount:            pricing.TaxAmount,
		Total:                pricing.Total,
		ShippingAddress:      req.ShippingAddress,
		ShippingPhoneNumber:  shippingPhone,
		BillingAddress:       req.BillingAddress,
		BillingPhoneNumber:   billingPhone,
		CustomerEmail:        req.CustomerEmail,
		CustomerPhone:        customerPhone,
		Notes:                req.Notes,
		IsGift:               isGift,
		GiftMessage:          req.GiftMessage,
		GiftCharge:           giftCharge,
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
	for _, item := range cartItems {
		product := products[item.ProductID]
		
		// Get seller-specific commission rate
		var seller models.User
		if err := tx.Select("commission").Where("id = ?", product.SellerID).First(&seller).Error; err != nil {
			return fmt.Errorf("failed to get seller commission: %w", err)
		}
		commissionRate := seller.Commission / 100.0 // Convert percentage to decimal
		
		// Use discount price if available, otherwise use regular price
		actualPrice := product.Price
		if product.DiscountPrice > 0 {
			actualPrice = product.DiscountPrice
		}
		
		itemTotal := actualPrice * float64(item.Quantity)
		commission := itemTotal * commissionRate
		sellerEarning := itemTotal - commission

		orderItem := &models.OrderItem{
			OrderID:        order.ID,
			ProductID:      item.ProductID,
			SellerID:       product.SellerID,
			ProductName:    product.Name,
			ProductSKU:     product.SKU,
			Quantity:       item.Quantity,
			Price:          actualPrice,
			Total:          itemTotal,
			Commission:     commission,
			SellerEarning:  sellerEarning,
			OriginalPrice:  product.Price,
			DiscountAmount: (product.Price - actualPrice) * float64(item.Quantity),
			Status:         string(models.OrderPending),
			IsGift:         item.IsGift,
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
	emailService := util.NewEmailService()
	// Send email to buyer
	if err := emailService.SendOrderPlacedEmail(order.CustomerEmail, order.OrderNumber, order.Total); err != nil {
		fmt.Printf("Failed to send order confirmation email to buyer: %v\n", err)
	}
	
	// Send email to admin
	adminEmail := os.Getenv("SUPER_ADMIN_EMAIL")
	if adminEmail != "" {
		if err := emailService.SendNewOrderNotificationToAdmin(adminEmail, order.OrderNumber, order.Total, order.CustomerEmail); err != nil {
			fmt.Printf("Failed to send order notification to admin: %v\n", err)
		}
	}
}

func (s *orderService) notifySellers(order *models.Order) {
	// Get order items with seller information
	var orderItems []models.OrderItem
	if err := s.db.Preload("Product").Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
		fmt.Printf("Failed to load order items: %v\n", err)
		return
	}
	
	// Group items by seller
	sellerItems := make(map[uint][]models.OrderItem)
	for _, item := range orderItems {
		sellerItems[item.SellerID] = append(sellerItems[item.SellerID], item)
	}
	
	// Send email to each seller
	emailService := util.NewEmailService()
	for sellerID, items := range sellerItems {
		// Get seller email
		var seller models.User
		if err := s.db.Preload("RegUser").First(&seller, sellerID).Error; err != nil {
			fmt.Printf("Failed to load seller %d: %v\n", sellerID, err)
			continue
		}
		
		if seller.RegUser == nil || seller.RegUser.Email == "" {
			fmt.Printf("Seller %d has no email\n", sellerID)
			continue
		}
		
		// Calculate total for this seller
		var itemTotal float64
		for _, item := range items {
			itemTotal += item.SellerEarning
		}
		
		if err := emailService.SendNewOrderNotificationToSeller(seller.RegUser.Email, order.OrderNumber, len(items), itemTotal); err != nil {
			fmt.Printf("Failed to send order notification to seller %d: %v\n", sellerID, err)
		}
	}
}

func (s *orderService) recordOrderMetrics(order *models.Order) {
}
