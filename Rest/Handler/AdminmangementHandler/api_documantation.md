# API Usage Examples

Complete guide for testing all API endpoints with cURL examples.

## 🔐 Authentication

### 1. Login as Super Admin

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@treestore.com",
    "password": "YourSecurePassword123"
  }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "admin@treestore.com",
    "role": "super_admin"
  }
}
```

**Note:** Copy the token and use it in subsequent requests.

---

## 👥 User Management

### 1. Get All Users

```bash
curl -X GET "http://localhost:8080/api/users?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Get Users by Status

```bash
# Get active users
curl -X GET "http://localhost:8080/api/users?page=1&limit=10&status=active" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get suspended users
curl -X GET "http://localhost:8080/api/users?page=1&limit=10&status=suspended" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Search Users

```bash
curl -X GET "http://localhost:8080/api/users/search?q=john&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Get User by ID

```bash
curl -X GET "http://localhost:8080/api/users/get?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get User Statistics

```bash
curl -X GET "http://localhost:8080/api/users/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 150,
    "active": 120,
    "inactive": 20,
    "suspended": 10
  }
}
```

### 6. Update User

```bash
curl -X PUT "http://localhost:8080/api/users/update?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+1234567890",
    "gender": "male"
  }'
```

### 7. Activate User

```bash
curl -X POST "http://localhost:8080/api/users/activate?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 8. Deactivate User

```bash
curl -X POST "http://localhost:8080/api/users/deactivate?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 9. Suspend User

```bash
curl -X POST "http://localhost:8080/api/users/suspend?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 🏪 Seller Management

### 1. Get All Sellers

```bash
curl -X GET "http://localhost:8080/api/sellers?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Get Sellers by Status

```bash
# Get approved sellers
curl -X GET "http://localhost:8080/api/sellers?status=approved&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get suspended sellers
curl -X GET "http://localhost:8080/api/sellers?status=suspended&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Get Pending Seller Approvals

```bash
curl -X GET "http://localhost:8080/api/sellers/pending?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Search Sellers

```bash
curl -X GET "http://localhost:8080/api/sellers/search?q=electronics&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get Seller by ID

```bash
curl -X GET "http://localhost:8080/api/sellers/get?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 6. Get Seller Statistics

```bash
curl -X GET "http://localhost:8080/api/sellers/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 50,
    "pending": 5,
    "approved": 40,
    "rejected": 3,
    "suspended": 2
  }
}
```

### 7. Get Top Sellers

```bash
curl -X GET "http://localhost:8080/api/sellers/top?limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 8. Approve Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/approve?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 9. Reject Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/reject?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Incomplete business documentation"
  }'
```

### 10. Suspend Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/suspend?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Multiple customer complaints"
  }'
```

### 11. Reactivate Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/reactivate?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 📦 Product Management

### 1. Get All Products

```bash
curl -X GET "http://localhost:8080/api/products?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Get Products by Status

```bash
# Get approved products
curl -X GET "http://localhost:8080/api/products?status=approved&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get pending products
curl -X GET "http://localhost:8080/api/products?status=pending&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Get Pending Product Approvals

```bash
curl -X GET "http://localhost:8080/api/products/pending?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Search Products

```bash
curl -X GET "http://localhost:8080/api/products/search?q=laptop&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get Product by ID

```bash
curl -X GET "http://localhost:8080/api/products/get?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 6. Get Product Statistics

```bash
curl -X GET "http://localhost:8080/api/products/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 500,
    "pending": 25,
    "approved": 450,
    "rejected": 15,
    "low_stock": 30
  }
}
```

### 7. Approve Product

```bash
curl -X POST "http://localhost:8080/api/products/approve?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 8. Reject Product

```bash
curl -X POST "http://localhost:8080/api/products/reject?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Product images do not match description"
  }'
```

### 9. Delete Product

```bash
curl -X DELETE "http://localhost:8080/api/products/delete?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 📋 Order Management

### 1. Get All Orders

```bash
curl -X GET "http://localhost:8080/api/orders?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Get Orders by Status

```bash
# Get pending orders
curl -X GET "http://localhost:8080/api/orders?status=pending&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get delivered orders
curl -X GET "http://localhost:8080/api/orders?status=delivered&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get cancelled orders
curl -X GET "http://localhost:8080/api/orders?status=cancelled&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Search Orders

```bash
# Search by order number
curl -X GET "http://localhost:8080/api/orders/search?q=ORD-12345&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Search by buyer email
curl -X GET "http://localhost:8080/api/orders/search?q=buyer@email.com&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Get Order by ID

```bash
curl -X GET "http://localhost:8080/api/orders/get?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get Order Statistics

```bash
curl -X GET "http://localhost:8080/api/orders/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 1500,
    "pending": 50,
    "confirmed": 100,
    "processing": 80,
    "shipped": 200,
    "delivered": 1000,
    "cancelled": 70,
    "total_revenue": 150000.50
  }
}
```
### update order stats now updated order stats give you montly 
```bash
curl -X GET "http://localhost:8080/api/orders/stats" \
  -H
  ```
  **response:**
```json
{
  "total": 120,
  "total_revenue": 54320.50,
  "statuses": {
    "pending": 15,
    "confirmed": 10,
    "processing": 20,
    "shipped": 30,
    "delivered": 40,
    "cancelled": 5
  },
  "monthly_report": [
    { "month": "2025-06", "revenue": 3200.00, "orders": 12 },
    { "month": "2025-07", "revenue": 5500.00, "orders": 20 },
    { "month": "2025-08", "revenue": 8800.00, "orders": 25 },
    { "month": "2025-09", "revenue": 6700.00, "orders": 22 },
    { "month": "2025-10", "revenue": 9500.00, "orders": 28 },
    { "month": "2025-11", "revenue": 7620.50, "orders": 13 }
  ]
}
```
### 6. Get Recent Orders

```bash
curl -X GET "http://localhost:8080/api/orders/recent?limit=20" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 7. Update Order Status

```bash
# Confirm order
curl -X PUT "http://localhost:8080/api/orders/update-status?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "confirmed"
  }'

# Mark as shipped
curl -X PUT "http://localhost:8080/api/orders/update-status?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "shipped"
  }'

# Mark as delivered
curl -X PUT "http://localhost:8080/api/orders/update-status?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "delivered"
  }'
```

### 8. Cancel Order

```bash
curl -X POST "http://localhost:8080/api/orders/cancel?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Customer requested cancellation"
  }'
```

---

## 📊 Dashboard

### Get Complete Dashboard Statistics

```bash
curl -X GET "http://localhost:8080/api/dashboard/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "users": {
      "total": 150,
      "active": 120,
      "inactive": 20,
      "suspended": 10
    },
    "sellers": {
      "total": 50,
      "pending": 5,
      "approved": 40,
      "rejected": 3,
      "suspended": 2
    },
    "products": {
      "total": 500,
      "pending": 25,
      "approved": 450,
      "rejected": 15,
      "low_stock": 30
    },
    "orders": {
      "total": 1500,
      "pending": 50,
      "confirmed": 100,
      "processing": 80,
      "shipped": 200,
      "delivered": 1000,
      "cancelled": 70,
      "total_revenue": 150000.50
    }
  }
}
```

---

## 🏥 Health Check

```bash
curl -X GET "http://localhost:8080/api/health"
```

**Response:**
```
OK
```

---

## 📝 Response Format

### Success Response

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

### Paginated Response

```json
{
  "success": true,
  "data": [ ... ],
  "page": 1,
  "limit": 10,
  "total": 100
}
```

### Error Response

```json
{
  "success": false,
  "error": "Error message here"
}
```

---

## 🔑 Common Status Codes

| Code | Meaning |
|------|---------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 500 | Internal Server Error |

---

## 💡 Tips

1. **Save your JWT token**: After login, save the token to use in all protected endpoints
2. **Pagination**: Use `page` and `limit` query parameters for large datasets
3. **Filtering**: Use `status` query parameter to filter results
4. **Search**: Use `q` query parameter for search functionality
5. **Error Handling**: Always check the response status code and error message

---

## 🧪 Testing Workflow

### Complete Testing Sequence

1. **Login**
   ```bash
   # Get JWT token
   curl -X POST http://localhost:8080/login ...
   ```

2. **Test User Management**
   ```bash
   # Get all users
   curl -X GET http://localhost:8080/api/users ...
   
   # Suspend a user
   curl -X POST http://localhost:8080/api/users/suspend?id=5 ...
   ```

3. **Test Seller Management**
   ```bash
   # Get pending sellers
   curl -X GET http://localhost:8080/api/sellers/pending ...
   
   # Approve a seller
   curl -X POST http://localhost:8080/api/sellers/approve?id=3 ...
   ```

4. **Test Product Management**
   ```bash
   # Get pending products
   curl -X GET http://localhost:8080/api/products/pending ...
   
   # Approve a product
   curl -X POST http://localhost:8080/api/products/approve?id=15 ...
   ```

5. **Test Order Management**
   ```bash
   # Get all orders
   curl -X GET http://localhost:8080/api/orders ...
   
   # Update order status
   curl -X PUT http://localhost:8080/api/orders/update-status?id=25 ...
   ```

6. **Check Dashboard**
   ```bash
   # Get complete statistics
   curl -X GET http://localhost:8080/api/dashboard/stats ...
   ```