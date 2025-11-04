package adminmangement

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// GetAll retrieves all orders with pagination
func (r *OrderRepository) GetAll(page, limit int, status string) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := r.db.Model(&models.Order{}).
		Preload("Buyer").
		Preload("Buyer.RegUser").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("OrderItems.Seller")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&orders).Error

	return orders, total, err
}

// GetByID retrieves an order by ID
func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Buyer").
		Preload("Buyer.RegUser").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("OrderItems.Seller").
		Preload("ShippingAddress").
		Preload("BillingAddress").
		Preload("Payment").
		First(&order, id).Error
	return &order, err
}

// GetByOrderNumber retrieves an order by order number
func (r *OrderRepository) GetByOrderNumber(orderNumber string) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Buyer").
		Preload("Buyer.RegUser").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("OrderItems.Seller").
		Preload("Payment").
		Where("order_number = ?", orderNumber).
		First(&order).Error
	return &order, err
}

// Update updates an order
func (r *OrderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

// UpdateStatus updates order status
func (r *OrderRepository) UpdateStatus(id uint, status string) error {
	updates := map[string]interface{}{"status": status}

	switch status {
	case "confirmed":
		updates["confirmed_at"] = time.Now()
	case "shipped":
		updates["shipped_at"] = time.Now()
	case "delivered":
		updates["delivered_at"] = time.Now()
	case "cancelled":
		updates["cancelled_at"] = time.Now()
	}

	return r.db.Model(&models.Order{}).Where("id = ?", id).Updates(updates).Error
}

// GetByBuyer retrieves orders by buyer ID
func (r *OrderRepository) GetByBuyer(buyerID uint, page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := r.db.Model(&models.Order{}).
		Where("buyer_id = ?", buyerID).
		Preload("OrderItems").
		Preload("OrderItems.Product")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&orders).Error

	return orders, total, err
}

// Search searches orders by order number or buyer email
func (r *OrderRepository) Search(query string, page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	searchQuery := "%" + query + "%"

	dbQuery := r.db.Model(&models.Order{}).
		Joins("LEFT JOIN buyers ON buyers.id = orders.buyer_id").
		Joins("LEFT JOIN reg_users ON reg_users.id = buyers.user_id").
		Where("orders.order_number ILIKE ? OR reg_users.email ILIKE ?", searchQuery, searchQuery).
		Preload("Buyer").
		Preload("Buyer.RegUser").
		Preload("OrderItems")

	// Count total results (use subquery for accuracy)
	if err := r.db.Table("(?) as sub", dbQuery.Select("orders.id")).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := dbQuery.Offset(offset).Limit(limit).Order("orders.created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// GetStats retrieves order statistics
// func (r *OrderRepository) GetStats() (map[string]interface{}, error) {
// 	stats := make(map[string]interface{})

// 	type Result struct {
// 		Status string
// 		Count  int64
// 	}

// 	var results []Result
// 	var totalRevenue float64
// 	var total int64

// 	// Group all statuses in one query
// 	if err := r.db.Model(&models.Order{}).
// 		Select("status, COUNT(*) as count").
// 		Group("status").
// 		Scan(&results).Error; err != nil {
// 		return nil, err
// 	}

// 	// Calculate total orders
// 	r.db.Model(&models.Order{}).Count(&total)

// 	//  Calculate total revenue
// 	if err := r.db.Model(&models.Order{}).
// 		Select("COALESCE(SUM(total), 0)").
// 		Where("status IN ?", []string{"delivered", "shipped", "processing"}).
// 		Scan(&totalRevenue).Error; err != nil {
// 		return nil, err
// 	}

// 	//  Initialize all status counts to 0
// 	stats["total"] = total
// 	stats["pending"] = int64(0)
// 	stats["confirmed"] = int64(0)
// 	stats["processing"] = int64(0)
// 	stats["shipped"] = int64(0)
// 	stats["delivered"] = int64(0)
// 	stats["cancelled"] = int64(0)

// 	//  Fill stats from query results
// 	for _, r := range results {
// 		stats[r.Status] = r.Count
// 	}

// 	stats["total_revenue"] = totalRevenue

// 	return stats, nil
// }

// montly revunue count with optimize version
func (r *OrderRepository) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	type Result struct {
		Status string
		Count  int64
	}

	var (
		results       []Result
		totalRevenue  float64
		totalOrders   int64
		monthlyReport []map[string]interface{}
	)

	// 1. Group all orders by status (fast single query)
	if err := r.db.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	//  2. Total order count
	if err := r.db.Model(&models.Order{}).Count(&totalOrders).Error; err != nil {
		return nil, err
	}

	// 3. Total revenue (for delivered, shipped, processing)
	if err := r.db.Model(&models.Order{}).
		Select("COALESCE(SUM(total), 0)").
		Where("status IN ?", []string{"delivered", "shipped", "processing"}).
		Scan(&totalRevenue).Error; err != nil {
		return nil, err
	}

	//  4. Monthly revenue + order count (last 6 months)
	if err := r.db.
		Model(&models.Order{}).
		Select(`
			TO_CHAR(DATE_TRUNC('month', created_at), 'YYYY-MM') AS month,
			COALESCE(SUM(total), 0) AS revenue,
			COUNT(*) AS orders
		`).
		Where("created_at >= NOW() - INTERVAL '6 months'").
		Group("month").
		Order("month ASC").
		Scan(&monthlyReport).Error; err != nil {
		return nil, err
	}

	//Initialize all status counts to 0
	statusMap := map[string]int64{
		"pending":    0,
		"confirmed":  0,
		"processing": 0,
		"shipped":    0,
		"delivered":  0,
		"cancelled":  0,
	}

	// ✅ Fill actual status counts
	for _, r := range results {
		statusMap[r.Status] = r.Count
	}

	// ✅ Combine results
	stats["total"] = totalOrders
	stats["total_revenue"] = totalRevenue
	stats["statuses"] = statusMap
	stats["monthly_report"] = monthlyReport

	return stats, nil
}

// GetRecentOrders retrieves recent orders
func (r *OrderRepository) GetRecentOrders(limit int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Buyer").
		Preload("Buyer.RegUser").
		Preload("OrderItems").
		Order("created_at DESC").
		Limit(limit).
		Find(&orders).Error
	return orders, err
}

// GetOrdersByDateRange retrieves orders within a date range
func (r *OrderRepository) GetOrdersByDateRange(startDate, endDate time.Time) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("OrderItems").
		Where("ordered_at BETWEEN ? AND ?", startDate, endDate).
		Order("ordered_at DESC").
		Find(&orders).Error
	return orders, err
}
