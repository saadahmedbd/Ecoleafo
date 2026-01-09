# Seller Dashboard Backend - Complete Documentation

## Table of Contents
1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Database Schema](#database-schema)
4. [API Endpoints](#api-endpoints)
5. [Implementation Guide](#implementation-guide)
6. [Code Structure](#code-structure)
7. [Testing](#testing)
8. [Troubleshooting](#troubleshooting)

---

## Overview

### Purpose
The Seller Dashboard Backend provides comprehensive analytics and statistics for sellers in a multi-vendor e-commerce platform. It aggregates data from orders, products, and reviews to present actionable insights.

### Technology Stack
- **Language**: Go (Golang) 1.21+
- **Framework**: net/http (standard library)
- **ORM**: GORM v2
- **Database**: PostgreSQL 14+
- **Authentication**: JWT Bearer tokens
- **Architecture**: Clean Architecture (Handler → Service → Repository)

### Key Features
- ✅ Real-time sales statistics
- ✅ Revenue analytics with charts
- ✅ Order status distribution
- ✅ Top selling products
- ✅ Low stock alerts
- ✅ Performance metrics
- ✅ Pending actions tracking
- ✅ Multi-seller support

---

## Architecture

### Layer Structure

```
┌─────────────────────────────────────────┐
│         HTTP Handler Layer              │
│  • Receives HTTP requests               │
│  • Validates JWT tokens                 │
│  • Extracts seller ID from context      │
│  • Calls service layer                  │
│  • Returns JSON responses               │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│         Service Layer                   │
│  • Validates input parameters           │
│  • Implements business logic            │
│  • Calls repository methods             │
│  • Handles errors                       │
│  • Returns DTOs                         │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│         Repository Layer                │
│  • Executes database queries            │
│  • Uses GORM for ORM                    │
│  • Aggregates data                      │
│  • Returns raw data                     │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│         PostgreSQL Database             │
│  • orders table                         │
│  • order_items table (seller_id FK)     │
│  • products table                       │
│  • users table (sellers)                │
└─────────────────────────────────────────┘
```

### Design Patterns

**1. Clean Architecture**
- Separation of concerns
- Dependency injection
- Testable code
- Independent layers

**2. Repository Pattern**
- Abstracts data access
- Centralizes queries
- Easy to test with mocks

**3. DTO Pattern**
- Type-safe API responses
- Aggregated data structures
- No model exposure

**4. Service Layer Pattern**
- Business logic isolation
- Reusable operations
- Validation centralization

---

## Database Schema

### Existing Tables (Used by Dashboard)

#### 1. orders
```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number VARCHAR(50) UNIQUE NOT NULL,
    buyer_id INTEGER NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    payment_status VARCHAR(20) NOT NULL DEFAULT 'pending',
    payment_method VARCHAR(50),
    subtotal DECIMAL(12,2) NOT NULL DEFAULT 0,
    shipping_cost DECIMAL(10,2) DEFAULT 0,
    tax_amount DECIMAL(10,2) DEFAULT 0,
    discount_amount DECIMAL(10,2) DEFAULT 0,
    total DECIMAL(12,2) NOT NULL,
    shipping_address TEXT NOT NULL DEFAULT 'n/a',
    billing_address TEXT,
    customer_email VARCHAR(100) NOT NULL,
    customer_phone VARCHAR(20),
    tracking_number VARCHAR(100),
    shipped_at TIMESTAMP,
    delivered_at TIMESTAMP,
    cancellation_reason TEXT,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    INDEX idx_buyer_id (buyer_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
);
```

**Status Values:**
- `pending` - Order created, awaiting confirmation
- `processing` - Seller confirmed, preparing items
- `shipped` - Order dispatched
- `delivered` - Order received by customer
- `cancelled` - Order cancelled

#### 2. order_items
```sql
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    seller_id INTEGER NOT NULL,        -- KEY: Tracks seller per item
    product_name VARCHAR(255) NOT NULL DEFAULT 'n/a',
    product_sku VARCHAR(100) NOT NULL DEFAULT 'n/a',
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price DECIMAL(10,2) NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    commission DECIMAL(10,2) NOT NULL,     -- Platform commission
    seller_earning DECIMAL(10,2) NOT NULL, -- Seller's earning after commission
    status VARCHAR(20) DEFAULT 'pending',
    shipped_at TIMESTAMP,
    delivered_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_order_id (order_id),
    INDEX idx_product_id (product_id),
    INDEX idx_seller_id (seller_id),      -- KEY: For seller queries
    FOREIGN KEY (order_id) REFERENCES orders(id)
);
```

**Important Fields:**
- `seller_id` - Links item to specific seller (multi-seller support)
- `seller_earning` - Actual earning after commission (pre-calculated)
- `commission` - Platform's cut (pre-calculated)

#### 3. products
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    seller_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(100) UNIQUE NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    discount_price DECIMAL(10,2),
    quantity INTEGER NOT NULL DEFAULT 0,
    min_quantity INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    is_approved BOOLEAN DEFAULT FALSE,
    is_featured BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    INDEX idx_seller_id (seller_id),
    INDEX idx_category_id (category_id),
    INDEX idx_is_active (is_active),
    INDEX idx_quantity (quantity)
);
```

#### 4. users (sellers)
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    business_email VARCHAR(100) UNIQUE,
    store_name VARCHAR(100) NOT NULL,
    store_slug VARCHAR(100) UNIQUE,
    total_sales DECIMAL(12,2) DEFAULT 0,
    total_earnings DECIMAL(12,2) DEFAULT 0,
    total_orders INTEGER DEFAULT 0,
    average_rating DECIMAL(3,2) DEFAULT 0,
    total_reviews INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    is_approved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Database Indexes for Performance

```sql
-- Order items by seller (critical for dashboard)
CREATE INDEX idx_order_items_seller_id ON order_items(seller_id);

-- Orders by status (for filtering)
CREATE INDEX idx_orders_status ON orders(status);

-- Orders by date (for analytics)
CREATE INDEX idx_orders_created_at ON orders(created_at);

-- Products by seller and quantity (for low stock)
CREATE INDEX idx_products_seller_quantity ON products(seller_id, quantity);

-- Composite index for common queries
CREATE INDEX idx_order_items_seller_order ON order_items(seller_id, order_id);
```

---

## API Endpoints

### Authentication
All endpoints require JWT Bearer token:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

The JWT token contains `user_id` which is used to identify the seller.

### Response Format
All endpoints return consistent JSON structure:

**Success Response:**
```json
{
  "status": "success",
  "message": "Statistics retrieved successfully",
  "data": { ... }
}
```

**Error Response:**
```json
{
  "status": "error",
  "message": "Invalid seller ID"
}
```

### Endpoint List

#### 1. Get Statistics
**GET** `/api/seller/statistics`

Returns comprehensive dashboard statistics.

**Request:**
```bash
curl -X GET http://localhost:8080/api/seller/statistics \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "total_sales": 45231.50,          // Sum of seller_earning
    "total_orders": 234,               // Distinct order count
    "total_earnings": 40708.35,        // From user.total_earnings
    "average_rating": 4.8,             // From user.average_rating
    "total_reviews": 89,               // From user.total_reviews
    "pending_orders": 23,              // Status = 'pending'
    "completed_orders": 195,           // Status = 'delivered'
    "shipped_orders": 12,              // Status = 'shipped'
    "cancelled_orders": 4,             // Status = 'cancelled'
    "active_products": 45,             // is_active = true
    "inactive_products": 5,            // is_active = false
    "low_stock_products": 8,           // quantity < 10
    "out_of_stock": 3,                 // quantity = 0
    "today_sales": 2340.50,            // Today's seller_earning
    "today_orders": 12,                // Today's order count
    "week_sales": 15420.75,            // Last 7 days
    "week_orders": 78,
    "month_sales": 45231.50,           // Last 30 days
    "month_orders": 234
  }
}
```

**Database Queries:**
```sql
-- Total sales (seller earnings)
SELECT COALESCE(SUM(seller_earning), 0) 
FROM order_items 
WHERE seller_id = ?

-- Total orders (distinct)
SELECT COUNT(DISTINCT order_id) 
FROM order_items 
WHERE seller_id = ?

-- Pending orders
SELECT COUNT(DISTINCT order_items.order_id)
FROM order_items
JOIN orders ON order_items.order_id = orders.id
WHERE order_items.seller_id = ? AND orders.status = 'pending'

-- Active products
SELECT COUNT(*) 
FROM products 
WHERE seller_id = ? AND is_active = true AND is_approved = true

-- Low stock products
SELECT COUNT(*) 
FROM products 
WHERE seller_id = ? AND quantity < 10 AND quantity > 0
```

#### 2. Get Sales Analytics
**GET** `/api/seller/analytics/sales?period=week`

Returns sales data grouped by time period for charts.

**Query Parameters:**
- `period` (string): `week`, `month`, or `year` (default: `week`)

**Response:**
```json
{
  "status": "success",
  "data": {
    "period": "week",
    "data": [
      {
        "name": "Mon",
        "date": "2024-11-04T00:00:00Z",
        "sales": 4200.50,
        "orders": 28,
        "revenue": 4200.50
      },
      {
        "name": "Tue",
        "date": "2024-11-05T00:00:00Z",
        "sales": 3800.25,
        "orders": 25,
        "revenue": 3800.25
      }
      // ... 5 more days
    ],
    "total": 29450.75,
    "change": 12.5  // Percentage change from previous period
  }
}
```

**Database Query:**
```sql
-- Week view (grouped by date)
SELECT 
    DATE(orders.created_at) as date,
    COALESCE(SUM(order_items.seller_earning), 0) as sales,
    COUNT(DISTINCT order_items.order_id) as orders
FROM order_items
JOIN orders ON order_items.order_id = orders.id
WHERE order_items.seller_id = ? 
  AND orders.created_at >= ?  -- Last 7 days
GROUP BY DATE(orders.created_at)
ORDER BY date ASC

-- Month view (grouped by date)
-- Same query with last 30 days

-- Year view (grouped by month)
SELECT 
    DATE_TRUNC('month', orders.created_at) as date,
    COALESCE(SUM(order_items.seller_earning), 0) as sales,
    COUNT(DISTINCT order_items.order_id) as orders
FROM order_items
JOIN orders ON order_items.order_id = orders.id
WHERE order_items.seller_id = ? 
  AND orders.created_at >= ?  -- Last 365 days
GROUP BY DATE_TRUNC('month', orders.created_at)
ORDER BY date ASC
```

#### 3. Get Revenue Analytics
**GET** `/api/seller/analytics/revenue?period=month`

Returns revenue breakdown with commission.

**Query Parameters:**
- `period` (string): `week`, `month`, or `year` (default: `month`)

**Response:**
```json
{
  "status": "success",
  "data": {
    "period": "month",
    "total_revenue": 53213.53,    // Gross (before commission)
    "net_revenue": 45231.50,      // seller_earning (after commission)
    "commission": 7982.03,        // Platform's cut
    "data": [
      {
        "name": "Week 1",
        "date": "2024-11-01T00:00:00Z",
        "revenue": 12500.00,
        "commission": 2206.25,
        "net": 10293.75
      }
      // ... more periods
    ],
    "change": 15.5
  }
}
```

**Calculation:**
```
seller_earning = item_total - commission  (stored in DB)
gross_revenue = seller_earning / (1 - commission_rate)
commission = gross_revenue - seller_earning
```

#### 4. Get Top Products
**GET** `/api/seller/products/top?limit=5`

Returns best-selling products.

**Query Parameters:**
- `limit` (int): Number of products (default: 5, max: 20)

**Response:**
```json
{
  "status": "success",
  "data": {
    "products": [
      {
        "id": 1,
        "name": "Oak Tree Sapling",
        "sku": "OAK-001",
        "price": 45.00,
        "total_sold": 89,
        "revenue": 4005.00,        // seller_earning for this product
        "image_url": "https://...",
        "stock": 23,
        "is_active": true,
        "is_approved": true
      }
      // ... 4 more products
    ],
    "count": 5
  }
}
```

**Database Query:**
```sql
SELECT 
    products.id,
    products.name,
    products.sku,
    products.price,
    COALESCE(SUM(order_items.quantity), 0) as total_sold,
    COALESCE(SUM(order_items.seller_earning), 0) as revenue,
    products.quantity as stock,
    products.is_active,
    products.is_approved
FROM products
LEFT JOIN order_items ON products.id = order_items.product_id
WHERE products.seller_id = ?
GROUP BY products.id
ORDER BY total_sold DESC
LIMIT ?
```

#### 5. Get Low Stock Products
**GET** `/api/seller/products/low-stock?threshold=10`

Returns products with low inventory.

**Query Parameters:**
- `threshold` (int): Stock threshold (default: 10)

**Response:**
```json
{
  "status": "success",
  "data": {
    "products": [
      {
        "id": 2,
        "name": "Pine Tree",
        "sku": "PINE-002",
        "quantity": 5,
        "min_stock": 10,
        "price": 65.00,
        "image_url": "https://...",
        "status": "low-stock",    // or "out-of-stock"
        "is_active": true
      }
      // ... more products
    ],
    "count": 8
  }
}
```

**Database Query:**
```sql
SELECT 
    id, name, sku, quantity, 
    min_quantity as min_stock, 
    price, is_active
FROM products
WHERE seller_id = ? AND quantity <= ?
ORDER BY quantity ASC
```

#### 6. Get Recent Orders
**GET** `/api/seller/orders/recent?limit=5`

Returns recent orders containing seller's products.

**Query Parameters:**
- `limit` (int): Number of orders (default: 5, max: 50)

**Response:**
```json
{
  "status": "success",
  "data": {
    "orders": [
      {
        "id": 123,
        "order_number": "ORD-1699012345-52",
        "customer_name": "John Doe",
        "product_name": "Oak Tree Sapling",
        "total": 45.00,              // seller_earning for this order
        "status": "pending",
        "created_at": "2024-11-02T10:30:00Z",
        "time_ago": "2 hours ago"
      }
      // ... 4 more orders
    ],
    "count": 5
  }
}
```

**Database Query:**
```sql
SELECT DISTINCT
    orders.id,
    orders.order_number,
    CONCAT(reg_users.first_name, ' ', reg_users.last_name) as customer_name,
    order_items.product_name,
    order_items.seller_earning as total,
    orders.status,
    orders.created_at
FROM orders
JOIN buyers ON orders.buyer_id = buyers.id
JOIN reg_users ON buyers.user_id = reg_users.id
JOIN order_items ON orders.id = order_items.order_id
WHERE order_items.seller_id = ?
ORDER BY orders.created_at DESC
LIMIT ?
```

#### 7. Get Order Distribution
**GET** `/api/seller/analytics/orders/distribution`

Returns order count by status for pie chart.

**Response:**
```json
{
  "status": "success",
  "data": {
    "distribution": [
      {
        "status": "delivered",
        "count": 195,
        "color": "#10B981",
        "value": 195
      },
      {
        "status": "pending",
        "count": 23,
        "color": "#F59E0B",
        "value": 23
      },
      {
        "status": "processing",
        "count": 10,
        "color": "#3B82F6",
        "value": 10
      },
      {
        "status": "shipped",
        "count": 12,
        "color": "#6366F1",
        "value": 12
      },
      {
        "status": "cancelled",
        "count": 4,
        "color": "#EF4444",
        "value": 4
      }
    ]
  }
}
```

**Database Query:**
```sql
SELECT 
    orders.status,
    COUNT(DISTINCT order_items.order_id) as count
FROM order_items
JOIN orders ON order_items.order_id = orders.id
WHERE order_items.seller_id = ?
GROUP BY orders.status
```

#### 8. Get Performance Metrics
**GET** `/api/seller/analytics/performance`

Returns key performance indicators.

**Response:**
```json
{
  "status": "success",
  "data": {
    "success_rate": 98.5,           // Delivered / Total orders
    "customer_satisfaction": 4.8,    // Average rating
    "average_order_value": 193.25,  // Total earnings / Total orders
    "fulfillment_rate": 95.2,       // (Shipped + Delivered) / Total
    "return_rate": 2.1,             // Cancelled / Total
    "conversion_rate": 2.5          // Placeholder
  }
}
```

**Calculations:**
```sql
-- Success rate
delivered_orders / total_orders * 100

-- Customer satisfaction
SELECT average_rating FROM users WHERE id = ?

-- Average order value
total_earnings / total_orders

-- Fulfillment rate
(shipped_orders + delivered_orders) / total_orders * 100

-- Return rate
cancelled_orders / total_orders * 100
```

#### 9. Get Pending Actions
**GET** `/api/seller/dashboard/pending-actions`

Returns items requiring seller attention.

**Response:**
```json
{
  "status": "success",
  "data": {
    "pending_orders": 23,      // Orders with status 'pending'
    "pending_products": 5,     // Products not approved
    "out_of_stock": 3,         // Products with quantity = 0
    "low_stock": 8,            // Products with quantity < 10
    "pending_reviews": 0,      // Future feature
    "unread_messages": 0       // Future feature
  }
}
```

**Database Queries:**
```sql
-- Pending orders
SELECT COUNT(DISTINCT order_items.order_id)
FROM order_items
JOIN orders ON order_items.order_id = orders.id
WHERE order_items.seller_id = ? AND orders.status = 'pending'

-- Pending products
SELECT COUNT(*) FROM products
WHERE seller_id = ? AND is_approved = false

-- Out of stock
SELECT COUNT(*) FROM products
WHERE seller_id = ? AND quantity = 0

-- Low stock
SELECT COUNT(*) FROM products
WHERE seller_id = ? AND quantity < 10 AND quantity > 0
```

---

## Implementation Guide

### Step 1: File Structure

```
backend/
├── dto/
│   └── dashboard_dto.go
├── repository/
│   └── dashboard_repository.go
├── service/
│   └── dashboard_service.go
├── handler/
│   └── dashboard_handler.go
├── routes/
│   └── dashboard_routes.go
└── main.go (update)
```

### Step 2: Create DTO File

**File:** `backend/dto/dashboard_dto.go`

Contains all response structures:
- `DashboardStatsResponse` - Main statistics
- `SalesAnalyticsResponse` - Chart data
- `SalesDataPoint` - Single data point
- `RecentOrderResponse` - Order list
- `TopProductResponse` - Product rankings
- `LowStockProductResponse` - Stock alerts
- `OrderStatusDistribution` - Pie chart data
- `PerformanceMetrics` - KPIs
- `PendingActionsResponse` - Action items
- `RevenueAnalyticsResponse` - Revenue breakdown
- `RevenueDataPoint` - Revenue data point

### Step 3: Create Repository File

**File:** `backend/repository/dashboard_repository.go`

Implements database queries:
- `GetSellerStatistics()` - Aggregate stats
- `GetSalesAnalytics()` - Time-series data
- `GetRecentOrders()` - Latest orders
- `GetTopProducts()` - Best sellers
- `GetLowStockProducts()` - Inventory alerts
- `GetOrderStatusDistribution()` - Status breakdown
- `GetPerformanceMetrics()` - KPI calculations
- `GetPendingActions()` - Action counts

**Key Points:**
- Uses GORM for ORM
- Joins order_items with orders
- Filters by `seller_id`
- Uses `seller_earning` field
- Handles NULL values with COALESCE
- Groups and aggregates data

### Step 4: Create Service File

**File:** `backend/service/dashboard_service.go`

Implements business logic:
- Validates seller ID
- Validates input parameters
- Calls repository methods
- Handles errors
- Returns DTOs

**Pattern:**
```go
func (s *DashboardService) GetStatistics(sellerID uint) (*dto.DashboardStatsResponse, error) {
    // 1. Validate input
    if sellerID == 0 {
        return nil, errors.New("invalid seller ID")
    }
    
    // 2. Call repository
    stats, err := s.dashboardRepo.GetSellerStatistics(sellerID)
    if err != nil {
        return nil, errors.New("failed to fetch statistics")
    }
    
    // 3. Return DTO
    return stats, nil
}
```

### Step 5: Create Handler File

**File:** `backend/handler/dashboard_handler.go`

Implements HTTP handlers:
- Extracts seller ID from JWT context
- Parses query parameters
- Calls service methods
- Returns JSON responses
- Handles HTTP errors

**Pattern:**
```go
func (h *DashboardHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
    // 1. Get seller ID from JWT
    sellerID, ok := r.Context().Value(middleware.UserIDKey).(uint)
    if !ok {
        utils.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // 2. Call service
    stats, err := h.dashboardService.GetSellerStatistics(sellerID)
    if err != nil {
        utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // 3. Return JSON
    utils.SuccessResponse(w, "Statistics retrieved successfully", stats, http.StatusOK)
}
```

### Step 6: Create Routes File

**File:** `backend/routes/dashboard_routes.go`

Maps URLs to handlers with middleware:
```go
func SetupDashboardRoutes(mux *http.ServeMux, dashboardHandler *handler.DashboardHandler) {
    mux.Handle("GET /api/seller/statistics", middleware.Chain(
        http.HandlerFunc(dashboardHandler.GetStatistics),
        middleware.Logger,
        middleware.Cors,
        middleware.AuthenticateJWT,
    ))
    
    // ... 8 more endpoints
}
```

### Step 7: Update main.go

Add dashboard initialization:
```go
// After existing repositories
dashboardRepo := repository.NewDashboardRepository(db)
dashboardService := service.NewDashboardService(dashboardRepo)
dashboardHandler := handler.NewDashboardHandler(dashboardService)

// After existing routes
routes.SetupDashboardRoutes(mux, dashboardHandler)
```

---

## Code Structure

### DTO Layer
**Purpose:** Define API response structures

**Example:**
```go
type DashboardStatsResponse struct {
    TotalSales    float64 `json:"total_sales"`
    TotalOrders   int     `json:"total_orders"`
    PendingOrders int     `json:"pending_orders"`
    // ... more fields
}
```

**Why DTOs:**
- Type safety
- API contract
- Aggregated data
- Hide sensitive fields
- Version control

### Repository Layer
**Purpose:** Database operations

**Example:**
```go
func (r *DashboardRepository) GetSellerStatistics(sellerID uint) (*dto.DashboardStatsResponse, error) {
    var stats dto.DashboardStatsResponse
    
    // Query 1: Get seller info
    var seller models.User
    r.db.Where("id = ?", sellerID).First(&seller)
    stats.AverageRating = seller.AverageRating
    
    // Query 2: Get order stats
    r.db.Table("order_items").
        Select("COUNT(DISTINCT order_id)").
        Where("seller_id = ?", sellerID).
        Scan(&stats.TotalOrders)
    
    // ... more queries
    
    return &stats, nil
}
```

**Best Practices:**
- Use GORM query builder
- Handle NULL values
- Use indexes
- Avoid N+1 queries
- Use joins efficiently

### Service Layer
**Purpose:** Business logic

**Example:**
```go
func (s *DashboardService) GetTopProducts(sellerID uint, limit int) ([]dto.TopProductResponse, error) {
    // Validate seller ID
    if sellerID == 0 {
        return nil, errors.New("invalid seller ID")
    }
    
    // Validate limit
    if limit <= 0 || limit > 20 {
        limit = 5
    }
    
    // Call repository
    products, err := s.dashboardRepo.GetTopProducts(sellerID, limit)
    if err != nil {
        return nil, errors.New("failed to fetch top products")
    }
    
    return products, nil
}
```

**Responsibilities:**
- Input validation
- Business rules
- Error handling
- Orchestration

### Handler Layer
**Purpose:** HTTP request/response

**Example:**
```go
func (h *DashboardHandler) GetTopProducts(w http.ResponseWriter, r *http.Request) {
    // Extract seller ID
    sellerID := r.Context().Value(middleware.UserIDKey).(uint)
    
    // Parse query params
    limitStr := r.URL.Query().Get("limit")
    limit := 5
    if limitStr != "" {
        limit, _ = strconv.Atoi(limitStr)
    }
    
    // Call service
    products, err := h.dashboardService.GetTopProducts(sellerID, limit)
    if err != nil {
        utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Return JSON
    utils.SuccessResponse(w, "Top products retrieved", map[string]interface{}{
        "products": products,
        "count": len(products),
    }, http.StatusOK)
}
```

**Responsibilities:**
- HTTP parsing
- Authentication check
- Response formatting
- Status codes

---

## Testing

### Unit Tests

**Repository Test:**
```go
func TestGetSellerStatistics(t *testing.T) {
    // Setup test database
    db := setupTestDB()
    repo := repository.NewDashboardRepository(db)
    
    // Insert test data
    sellerID := createTestSeller(db)
    createTestOrders(db, sellerID, 5)
    
    // Test
    stats, err := repo.GetSellerStatistics(sellerID)
    
    // Assertions
    assert.NoError(t, err)
    assert.Equal(t, 5, stats.TotalOrders)
    assert.Greater(t, stats.TotalSales, 0.0)
}
```

**Service Test:**
```go
func TestGetStatistics_InvalidSellerID(t *testing.T) {
    mockRepo := &MockDashboardRepository{}
    service := service.NewDashboardService(mockRepo)
    
    // Test with invalid ID
    stats, err := service.GetSellerStatistics(0)
    
    // Assertions
    assert.Error(t, err)
    assert.Nil(t, stats)
    assert.Equal(t, "invalid seller ID", err.Error())
}
```

### Integration Tests

**API Test:**
```bash
# Test statistics endpoint
curl -X GET http://localhost:8080/api/seller/statistics \
  -H "Authorization: Bearer $(cat token.txt)" \
  -H "Content-Type: application/json"

# Expected: 200 OK with statistics JSON
```

**Load Test with Apache Bench:**
```bash
ab -n 1000 -c 10 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  http://localhost:8080/api/seller/statistics
```

### Test Data Setup

```sql
-- Create test seller
INSERT INTO users (id, store_name, is_active, is_approved) 
VALUES (999, 'Test Store', true, true);

-- Create test orders
INSERT INTO orders (id, order_number, buyer_id, status, total, created_at)
VALUES 
  (1001, 'ORD-TEST-001', 1, 'pending', 100.00, NOW() - INTERVAL '1 day'),
  (1002, 'ORD-TEST-002', 1, 'delivered', 200.00, NOW() - INTERVAL '2 days'),
  (1003, 'ORD-TEST-003', 1, 'shipped', 150.00, NOW() - INTERVAL '3 days');

-- Create test order items
INSERT INTO order_items (order_id, product_id, seller_id, quantity, price, total, commission, seller_earning)
VALUES 
  (1001, 1, 999, 2, 50.00, 100.00, 15.00, 85.00),
  (1002, 2, 999, 1, 200.00, 200.00, 30.00, 170.00),
  (1003, 3, 999, 3, 50.00, 150.00, 22.50, 127.50);

-- Create test products
INSERT INTO products (id, seller_id, name, sku, price, quantity, is_active, is_approved)
VALUES 
  (1, 999, 'Test Product 1', 'TEST-001', 50.00, 100, true, true),
  (2, 999, 'Test Product 2', 'TEST-002', 200.00, 5, true, true),
  (3, 999, 'Test Product 3', 'TEST-003', 50.00, 0, true, true);
```

---

## Troubleshooting

### Common Issues

#### 1. Statistics Return Zero

**Symptom:** All statistics show 0

**Causes:**
- No orders for seller
- `seller_id` mismatch in order_items
- JWT token contains wrong user_id

**Solution:**
```sql
-- Check if seller has orders
SELECT COUNT(*) FROM order_items WHERE seller_id = YOUR_SELLER_ID;

-- Check seller_id in JWT
-- Print in handler: log.Printf("Seller ID from JWT: %d", sellerID)

-- Verify order_items have correct seller_id
SELECT DISTINCT seller_id FROM order_items;
```

#### 2. JWT Unauthorized Error

**Symptom:** 401 Unauthorized

**Causes:**
- Missing Authorization header
- Invalid token format
- Expired token
- Wrong JWT secret

**Solution:**
```go
// Verify token format
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

// Check JWT secret matches
JWT_SECRET=your-secret-key (must be same everywhere)

// Check token expiration
// Decode JWT at jwt.io to see exp claim
```

#### 3. Database Query Errors

**Symptom:** "failed to fetch statistics"

**Causes:**
- Missing indexes
- Column name mismatch
- Table doesn't exist
- Foreign key issues

**Solution:**
```sql
-- Verify tables exist
\dt

-- Check columns match code
\d order_items
\d orders
\d products

-- Add missing indexes
CREATE INDEX idx_order_items_seller_id ON order_items(seller_id);
CREATE INDEX idx_orders_status ON orders(status);
```

#### 4. Type Assertion Panic

**Symptom:** `panic: interface conversion: interface {} is nil`

**Causes:**
- seller_id not in JWT context
- Middleware not applied
- Context value wrong type

**Solution:**
```go
// Safe type assertion
sellerID, ok := r.Context().Value(middleware.UserIDKey).(uint)
if !ok {
    utils.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
    return
}

// Verify middleware is applied in routes
middleware.Chain(
    http.HandlerFunc(handler.GetStatistics),
    middleware.AuthenticateJWT,  // Must be here!
)
```

#### 5. CORS Errors

**Symptom:** CORS policy blocked in browser

**Causes:**
- Missing CORS headers
- Wrong allowed origin
- Preflight request fails

**Solution:**
```go
// In middleware/cors.go
func Cors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

#### 6. Slow Queries

**Symptom:** API responses take > 1 second

**Causes:**
- Missing indexes
- N+1 query problem
- Large dataset without pagination
- Inefficient joins

**Solution:**
```sql
-- Add indexes
CREATE INDEX idx_order_items_seller_order ON order_items(seller_id, order_id);
CREATE INDEX idx_orders_created_at ON orders(created_at);

-- Use EXPLAIN to analyze queries
EXPLAIN ANALYZE
SELECT COUNT(DISTINCT order_id) 
FROM order_items 
WHERE seller_id = 999;

-- Should show "Index Scan" not "Seq Scan"
```

#### 7. Memory Leaks

**Symptom:** Server memory increases over time

**Causes:**
- Unclosed database connections
- Goroutine leaks
- Caching without limits

**Solution:**
```go
// Always use connection pooling
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxLifetime(5 * time.Minute)

// Close result sets if using raw SQL
rows, err := db.Query("...")
defer rows.Close()

// Use context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

### Debug Mode

Enable debug logging in GORM:
```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
```

This will log all SQL queries to console.

### Performance Monitoring

```go
// Add request timing middleware
func TimingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start)
        log.Printf("%s %s took %v", r.Method, r.URL.Path, duration)
    })
}
```

---

## Security Considerations

### 1. SQL Injection Prevention
GORM automatically escapes parameters:
```go
// Safe - GORM parameterizes
db.Where("seller_id = ?", sellerID).Find(&orders)

// Unsafe - Don't do this
db.Raw("SELECT * FROM orders WHERE seller_id = " + sellerID)
```

### 2. Authorization
Always verify seller owns the data:
```go
// Get seller ID from JWT (trusted source)
sellerID := r.Context().Value(middleware.UserIDKey).(uint)

// Query only that seller's data
db.Where("seller_id = ?", sellerID).Find(&products)
```

### 3. Rate Limiting
Implement rate limiting for APIs:
```go
// Example with golang.org/x/time/rate
limiter := rate.NewLimiter(rate.Limit(100), 200) // 100 req/sec, burst 200

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### 4. Input Validation
Validate all user inputs:
```go
func (s *DashboardService) GetSalesAnalytics(sellerID uint, period string) error {
    // Validate seller ID
    if sellerID == 0 {
        return errors.New("invalid seller ID")
    }
    
    // Whitelist validation
    validPeriods := map[string]bool{"week": true, "month": true, "year": true}
    if !validPeriods[period] {
        return errors.New("invalid period")
    }
    
    // Range validation
    if limit < 1 || limit > 100 {
        return errors.New("limit must be between 1 and 100")
    }
}
```

### 5. Error Handling
Don't expose internal errors:
```go
// Bad - exposes DB structure
return errors.New("pq: relation order_items does not exist")

// Good - generic error
return errors.New("failed to fetch statistics")
```

---

## Performance Optimization

### 1. Database Indexing
```sql
-- Critical indexes for dashboard queries
CREATE INDEX idx_order_items_seller_id ON order_items(seller_id);
CREATE INDEX idx_order_items_seller_order ON order_items(seller_id, order_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_created_at ON orders(created_at);
CREATE INDEX idx_products_seller_quantity ON products(seller_id, quantity);

-- Composite index for common query patterns
CREATE INDEX idx_orders_seller_status_date ON orders(seller_id, status, created_at);
```

### 2. Query Optimization
```go
// Bad - N+1 query problem
for _, product := range products {
    var image ProductImage
    db.Where("product_id = ?", product.ID).First(&image)
}

// Good - Single query with join
db.Preload("Images").Find(&products)

// Better - Specific preload
db.Preload("Images", func(db *gorm.DB) *gorm.DB {
    return db.Order("is_primary DESC, sort_order ASC").Limit(1)
}).Find(&products)
```

### 3. Caching Strategy
```go
// Use Redis for frequently accessed data
type DashboardService struct {
    repo  *repository.DashboardRepository
    cache *redis.Client
}

func (s *DashboardService) GetStatistics(sellerID uint) (*dto.DashboardStatsResponse, error) {
    // Try cache first
    cacheKey := fmt.Sprintf("stats:seller:%d", sellerID)
    
    var stats dto.DashboardStatsResponse
    err := s.cache.Get(ctx, cacheKey).Scan(&stats)
    if err == nil {
        return &stats, nil // Cache hit
    }
    
    // Cache miss - fetch from DB
    stats, err := s.repo.GetSellerStatistics(sellerID)
    if err != nil {
        return nil, err
    }
    
    // Cache for 5 minutes
    s.cache.Set(ctx, cacheKey, stats, 5*time.Minute)
    
    return stats, nil
}
```

### 4. Connection Pooling
```go
// Configure database connection pool
db.SetMaxOpenConns(25)          // Max open connections
db.SetMaxIdleConns(25)          // Max idle connections
db.SetConnMaxLifetime(5 * time.Minute)  // Connection lifetime
db.SetConnMaxIdleTime(10 * time.Minute) // Idle timeout
```

### 5. Pagination
Always paginate large result sets:
```go
func (r *DashboardRepository) GetRecentOrders(sellerID uint, limit int, offset int) ([]dto.RecentOrderResponse, error) {
    var orders []dto.RecentOrderResponse
    
    err := r.db.
        Limit(limit).
        Offset(offset).
        Order("created_at DESC").
        Find(&orders).Error
    
    return orders, err
}
```

---

## Deployment Checklist

### Pre-Deployment

- [ ] All tests pass
- [ ] Database indexes created
- [ ] Environment variables configured
- [ ] JWT secret is strong and unique
- [ ] CORS origins set correctly
- [ ] Rate limiting enabled
- [ ] Logging configured
- [ ] Error handling tested
- [ ] Security audit done

### Production Configuration

```env
# Production .env
DB_HOST=production-db.example.com
DB_PORT=5432
DB_USER=prod_user
DB_PASSWORD=strong_random_password_here
DB_NAME=ecommerce_prod

SERVER_PORT=8080

JWT_SECRET=very-strong-random-secret-key-at-least-32-characters
JWT_EXPIRATION=24h

# Cloudinary
CLOUDINARY_CLOUD_NAME=prod_cloud
CLOUDINARY_API_KEY=prod_key
CLOUDINARY_API_SECRET=prod_secret

# Performance
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=25
DB_CONN_MAX_LIFETIME=5m

# Security
ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_BURST=200
```

### Monitoring

```go
// Add health check endpoint
func (h *DashboardHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
    // Check database connection
    err := h.dashboardService.repo.db.Exec("SELECT 1").Error
    if err != nil {
        w.WriteHeader(http.StatusServiceUnavailable)
        json.NewEncoder(w).Encode(map[string]string{
            "status": "unhealthy",
            "error": "database connection failed",
        })
        return
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}
```

### Backup Strategy

```bash
# Daily database backup
0 2 * * * pg_dump -h localhost -U postgres ecommerce_prod > /backups/db_$(date +\%Y\%m\%d).sql

# Keep last 7 days
find /backups -name "db_*.sql" -mtime +7 -delete
```

---

## API Response Examples

### Success Response Format
```json
{
  "status": "success",
  "message": "Statistics retrieved successfully",
  "data": {
    "total_sales": 45231.50,
    "total_orders": 234
  }
}
```

### Error Response Format
```json
{
  "status": "error",
  "message": "Invalid seller ID"
}
```

### Validation Error Format
```json
{
  "status": "error",
  "message": "Validation failed",
  "errors": {
    "period": "must be one of: week, month, year",
    "limit": "must be between 1 and 100"
  }
}
```

---

## Version History

### v1.0.0 (Current)
- Initial release
- 9 API endpoints
- Multi-seller support
- Real-time statistics
- Performance optimized

### Future Enhancements
- v1.1.0: Add date range filtering
- v1.2.0: Export to PDF/CSV
- v1.3.0: Real-time WebSocket updates
- v1.4.0: Predictive analytics
- v1.5.0: Custom reports

---

## Support & Resources

### Internal Documentation
- API Endpoints: `/docs/api`
- Database Schema: `/docs/database`
- Frontend Integration: `/docs/frontend`

### External Resources
- GORM Documentation: https://gorm.io/docs/
- Go net/http: https://pkg.go.dev/net/http
- PostgreSQL: https://www.postgresql.org/docs/

### Getting Help
1. Check this documentation
2. Review error logs
3. Test endpoints with curl
4. Check database connections
5. Verify JWT tokens

---

**Document Version**: 1.0.0  
**Last Updated**: November 2024  
**Status**: Production Ready ✅  
**Maintained By**: Development Team