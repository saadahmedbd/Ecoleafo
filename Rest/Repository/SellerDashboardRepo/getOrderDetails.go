package sellerdashboardrepo

import (
	"errors"

	sellerdashboard "github.com/saadahmedbd/Treestore/Rest/DTO/Seller_Dashboard"
)

// GetOrderDetails - Get order details by ID
func (r *DashboardRepository) GetOrderDetails(regUserID uint, orderID uint) (*sellerdashboard.OrderDetailsResponse, error) {
	// Convert reguser ID to seller ID
	var seller struct {
		ID uint
	}
	if err := r.db.Table("users").Select("id").Where("user_id = ?", regUserID).First(&seller).Error; err != nil {
		return nil, err
	}
	sellerID := seller.ID

	var order sellerdashboard.OrderDetailsResponse

	// Get order details
	err := r.db.Table("orders").
		Select(`orders.id, orders.order_number, orders.status, orders.payment_status, 
			orders.payment_method, orders.customer_email, orders.customer_phone, 
			orders.shipping_address, orders.created_at,
			CONCAT(reg_users.first_name, ' ', reg_users.last_name) as customer_name`).
		Joins("JOIN buyers ON orders.buyer_id = buyers.id").
		Joins("JOIN reg_users ON buyers.user_id = reg_users.id").
		Where("orders.id = ?", orderID).
		Scan(&order).Error

	if err != nil {
		return nil, err
	}

	// Verify this order has items from this seller
	var itemCount int64
	r.db.Table("order_items").
		Where("order_id = ? AND seller_id = ?", orderID, sellerID).
		Count(&itemCount)

	if itemCount == 0 {
		return nil, errors.New("order not found or does not belong to this seller")
	}

	// Get order items for this seller
	var items []sellerdashboard.OrderItemDetails
	r.db.Table("order_items").
		Select("id, product_id, product_name, product_sku, quantity, price, total, commission, seller_earning").
		Where("order_id = ? AND seller_id = ?", orderID, sellerID).
		Scan(&items)

	order.Items = items

	// Calculate totals for this seller's items
	var totals struct {
		Total      float64
		Commission float64
		NetEarning float64
	}
	r.db.Table("order_items").
		Select("COALESCE(SUM(total), 0) as total, COALESCE(SUM(commission), 0) as commission, COALESCE(SUM(seller_earning), 0) as net_earning").
		Where("order_id = ? AND seller_id = ?", orderID, sellerID).
		Scan(&totals)

	order.Total = totals.Total
	order.Commission = totals.Commission
	order.NetEarning = totals.NetEarning

	return &order, nil
}
