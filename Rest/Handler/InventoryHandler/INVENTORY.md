# Seller Inventory API Documentation

## Overview
The Seller Inventory API provides endpoints for sellers to manage their product inventory, track stock levels, view history, and perform bulk operations.

## Base URL
`/api/seller/inventory`

## Authentication
All endpoints require JWT authentication with seller role.

---

## Endpoints

### 1. Get Inventory
**GET** `/api/seller/inventory`

Retrieve paginated inventory list with filtering options.

**Query Parameters:**
- `search` (string, optional) - Search by product name or SKU
- `status` (string, optional) - Filter by stock status: `all`, `low-stock`, `out-of-stock`, `in-stock`
- `category` (string, optional) - Filter by category ID
- `page` (int, optional, default: 1) - Page number
- `limit` (int, optional, default: 50, max: 100) - Items per page

**Response:**
```json
{
  "status": "success",
  "message": "Inventory retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "name": "Product Name",
        "sku": "SKU123",
        "category_id": 5,
        "category_name": "Electronics",
        "price": 99.99,
        "stock": 50,
        "low_stock_threshold": 10,
        "image_url": "https://...",
        "is_active": true,
        "is_approved": true,
        "last_updated": "2024-01-15T10:30:00Z",
        "stock_value": 4999.50
      }
    ],
    "total": 100,
    "page": 1,
    "limit": 50
  }
}
```

---

### 2. Get Inventory Stats
**GET** `/api/seller/inventory/stats`

Get inventory statistics and metrics.

**Response:**
```json
{
  "status": "success",
  "message": "Stats retrieved successfully",
  "data": {
    "total_products": 100,
    "low_stock_count": 15,
    "out_of_stock_count": 5,
    "total_stock_value": 50000.00
  }
}
```

---

### 3. Update Stock
**PUT** `/api/seller/inventory/{productId}/stock`

Update stock quantity for a single product.

**Path Parameters:**
- `productId` (uint, required) - Product ID

**Request Body:**
```json
{
  "quantity": 100
}
```

**Response:**
```json
{
  "status": "success",
  "message": "stock update successfully",
  "data": {
    "id": 1,
    "name": "Product Name",
    "stock": 100,
    "stock_value": 9999.00
  }
}
```

---

### 4. Bulk Update Stock
**POST** `/api/seller/inventory/bulk-update`

Update stock for multiple products at once (max 100 products).

**Request Body:**
```json
{
  "updates": [
    {
      "product_id": 1,
      "quantity": 50
    },
    {
      "product_id": 2,
      "quantity": 75
    }
  ]
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Bulk update successfully",
  "data": {
    "success_count": 2,
    "failed_count": 0,
    "errors": []
  }
}
```

---

### 5. Update Low Stock Threshold
**PUT** `/api/seller/inventory/{productId}/threshold`

Update the low stock alert threshold for a product.

**Path Parameters:**
- `productId` (uint, required) - Product ID

**Request Body:**
```json
{
  "min_quantity": 10
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Threshold updated successfully",
  "data": {
    "id": 1,
    "low_stock_threshold": 10
  }
}
```

---

### 6. Get Stock History
**GET** `/api/seller/inventory/{productId}/history`

Get stock change history for a product.

**Path Parameters:**
- `productId` (uint, required) - Product ID

**Query Parameters:**
- `limit` (int, optional, default: 50) - Number of history records

**Response:**
```json
{
  "status": "success",
  "message": "History retrieved successfully",
  "data": [
    {
      "id": 1,
      "product_id": 1,
      "type": "adjustment",
      "quantity": 10,
      "prev_stock": 40,
      "new_stock": 50,
      "reason": "bulk_update",
      "reference": "Bulk stock update",
      "created_by": 5,
      "created_at": "2024-01-15T10:30:00Z"
    }
  ]
}
```

---

### 7. Export Inventory
**GET** `/api/seller/inventory/export`

Export inventory data (CSV/Excel format).

**Query Parameters:**
- `format` (string, optional, default: csv) - Export format: `csv`, `excel`

**Response:**
File download with inventory data.

---

## Error Responses

**401 Unauthorized:**
```json
{
  "error": "unauthorized"
}
```

**400 Bad Request:**
```json
{
  "status": "error",
  "message": "Invalid request body",
  "data": null
}
```

**500 Internal Server Error:**
```json
{
  "status": "error",
  "message": "Failed to fetch inventory",
  "data": null
}
```

---

## Stock Status Types
- `all` - All products
- `in-stock` - Products with stock above threshold
- `low-stock` - Products with stock at or below threshold but not zero
- `out-of-stock` - Products with zero stock

## Stock Change Types
- `adjustment` - Manual stock adjustment
- `restock` - Stock increase
- `sale` - Stock decrease from order
- `bulk_update` - Bulk update operation
- `manual` - Manual update

---

## Notes
- All timestamps are in ISO 8601 format (UTC)
- Stock values are calculated as `quantity * price`
- Bulk updates are limited to 100 products per request
- History records are ordered by most recent first
